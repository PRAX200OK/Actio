package mcp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"sync"
	"sync/atomic"
	"time"
)

type pluginClient interface {
	ListResources(ctx context.Context) ([]map[string]any, error)
	ReadResource(ctx context.Context, uri string) (string, error)
	Close() error
}

type stdioPluginClient struct {
	cfg pluginConfig

	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser

	nextID uint64

	mu      sync.Mutex
	pending map[string]chan response

	closeOnce sync.Once
	closed    chan struct{}
}

func newStdioPluginClient(ctx context.Context, cfg pluginConfig) (*stdioPluginClient, error) {
	cmd := exec.CommandContext(ctx, cfg.Command, cfg.Args...)
	if cfg.Env != nil {
		for k, v := range cfg.Env {
			cmd.Env = append(cmd.Env, k+"="+v)
		}
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("stdin pipe: %w", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		_ = stdin.Close()
		return nil, fmt.Errorf("stdout pipe: %w", err)
	}

	// Start the plugin server process.
	if err := cmd.Start(); err != nil {
		_ = stdin.Close()
		_ = stdout.Close()
		return nil, fmt.Errorf("start plugin %q: %w", cfg.Name, err)
	}

	c := &stdioPluginClient{
		cfg:     cfg,
		cmd:     cmd,
		stdin:   stdin,
		stdout:  stdout,
		pending: make(map[string]chan response),
		closed:  make(chan struct{}),
	}

	go c.readLoop()
	return c, nil
}

func (c *stdioPluginClient) Close() error {
	var err error
	c.closeOnce.Do(func() {
		close(c.closed)
		_ = c.stdin.Close()
		_ = c.stdout.Close()
		if c.cmd.Process != nil {
			_ = c.cmd.Process.Kill()
		}
		_, err = c.cmd.Process.Wait()
	})
	return err
}

func (c *stdioPluginClient) readLoop() {
	sc := bufio.NewScanner(c.stdout)
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 {
			continue
		}
		var resp response
		if err := json.Unmarshal(line, &resp); err != nil {
			continue
		}

		idKey := string(resp.ID)
		c.mu.Lock()
		ch := c.pending[idKey]
		if ch != nil {
			delete(c.pending, idKey)
		}
		c.mu.Unlock()
		if ch != nil {
			ch <- resp
			close(ch)
		}
	}

	// If stdout closes, fail any pending requests quickly.
	c.mu.Lock()
	for id, ch := range c.pending {
		delete(c.pending, id)
		close(ch)
	}
	c.mu.Unlock()
}

func (c *stdioPluginClient) call(ctx context.Context, method string, params any) (response, error) {
	id := atomic.AddUint64(&c.nextID, 1)
	idRaw := json.RawMessage([]byte(fmt.Sprintf("%d", id)))
	idKey := string(idRaw)

	var paramsRaw json.RawMessage
	if params != nil {
		b, err := json.Marshal(params)
		if err != nil {
			return response{}, err
		}
		paramsRaw = b
	} else {
		paramsRaw = json.RawMessage([]byte(`{}`))
	}

	req := request{
		JsonRPC: "2.0",
		ID:      idRaw,
		Method:  method,
		Params:  paramsRaw,
	}
	data, err := json.Marshal(req)
	if err != nil {
		return response{}, err
	}

	ch := make(chan response, 1)
	c.mu.Lock()
	c.pending[idKey] = ch
	c.mu.Unlock()

	// Write request as a single line (newline-delimited JSON).
	if _, err := c.stdin.Write(append(data, '\n')); err != nil {
		c.mu.Lock()
		delete(c.pending, idKey)
		c.mu.Unlock()
		return response{}, fmt.Errorf("write request: %w", err)
	}

	// Wait for response or context cancellation.
	select {
	case <-ctx.Done():
		c.mu.Lock()
		delete(c.pending, idKey)
		c.mu.Unlock()
		return response{}, ctx.Err()
	case <-c.closed:
		return response{}, fmt.Errorf("plugin client closed")
	case resp, ok := <-ch:
		if !ok {
			return response{}, fmt.Errorf("plugin process closed")
		}
		return resp, nil
	case <-time.After(10 * time.Second):
		c.mu.Lock()
		delete(c.pending, idKey)
		c.mu.Unlock()
		return response{}, fmt.Errorf("plugin call timed out")
	}
}

func (c *stdioPluginClient) ListResources(ctx context.Context) ([]map[string]any, error) {
	resp, err := c.call(ctx, "mcp.listResources", map[string]any{})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf("plugin error: %s", resp.Error.Message)
	}
	m, _ := resp.Result.(map[string]any)
	raw, _ := m["resources"]
	res, _ := raw.([]map[string]any)
	if res == nil {
		// json.Unmarshal decodes []interface{} by default
		if a, ok := raw.([]any); ok {
			var out []map[string]any
			for _, it := range a {
				if mm, ok := it.(map[string]any); ok {
					out = append(out, mm)
				}
			}
			return out, nil
		}
		return []map[string]any{}, nil
	}
	return res, nil
}

func (c *stdioPluginClient) ReadResource(ctx context.Context, uri string) (string, error) {
	resp, err := c.call(ctx, "mcp.readResource", map[string]string{"uri": uri})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf("plugin error: %s", resp.Error.Message)
	}
	m, _ := resp.Result.(map[string]any)
	if s, ok := m["content"].(string); ok {
		return s, nil
	}
	return "", fmt.Errorf("unexpected plugin response")
}

