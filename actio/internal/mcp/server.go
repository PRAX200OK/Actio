package mcp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"actio/internal/actio"
)

// Basic JSON-RPC 2.0 types for a minimal MCP-style server.

type request struct {
	JsonRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type response struct {
	JsonRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Result  any             `json:"result,omitempty"`
	Error   *respError      `json:"error,omitempty"`
}

type respError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type server struct {
	root      string
	stderr    io.Writer
	pluginMgr *pluginManager
}

// ServeStdIO starts a minimal MCP-style JSON-RPC server over stdio.
// It intentionally supports a narrow tool surface focused on Actio context:
// - mcp.listResources
// - mcp.readResource
func ServeStdIO(ctx context.Context, root string, stderr io.Writer) error {
	s := &server{
		root:      root,
		stderr:    stderr,
		pluginMgr: newPluginManager(root, stderr, nil),
	}
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		line := in.Bytes()
		if len(line) == 0 {
			continue
		}

		var req request
		if err := json.Unmarshal(line, &req); err != nil {
			fmt.Fprintf(stderr, "invalid JSON-RPC request: %v\n", err)
			continue
		}

		resp := s.handleRequest(ctx, req)
		data, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprintf(stderr, "failed to encode response: %v\n", err)
			continue
		}
		fmt.Println(string(data))
	}

	if err := in.Err(); err != nil {
		return fmt.Errorf("mcp server input error: %w", err)
	}
	return nil
}

func (s *server) handleRequest(ctx context.Context, req request) response {
	switch req.Method {
	case "mcp.listResources":
		actioRes := listResources(s.root)
		pluginRes, _ := s.pluginMgr.listAllResources(ctx)
		return response{
			JsonRPC: "2.0",
			ID:      req.ID,
			Result: map[string]any{
				"resources": append(actioRes, pluginRes...),
			},
		}
	case "mcp.readResource":
		var params struct {
			URI string `json:"uri"`
		}
		if err := json.Unmarshal(req.Params, &params); err != nil {
			return errorResponse(req.ID, -32602, "invalid params")
		}
		var content string
		if pluginName, origURI, ok := parsePluginURI(params.URI); ok {
			c, err := s.pluginMgr.getClient(ctx, pluginName)
			if err != nil {
				return errorResponse(req.ID, -32001, err.Error())
			}
			content, err = c.ReadResource(ctx, origURI)
			if err != nil {
				return errorResponse(req.ID, -32001, err.Error())
			}
		} else {
			var err error
			content, err = readResource(s.root, params.URI)
			if err != nil {
				return errorResponse(req.ID, -32001, err.Error())
			}
		}
		return response{
			JsonRPC: "2.0",
			ID:      req.ID,
			Result: map[string]any{
				"uri":     params.URI,
				"content": content,
			},
		}
	default:
		return errorResponse(req.ID, -32601, "method not found")
	}
}

func errorResponse(id json.RawMessage, code int, msg string) response {
	return response{
		JsonRPC: "2.0",
		ID:      id,
		Error: &respError{
			Code:    code,
			Message: msg,
		},
	}
}

func listResources(root string) []map[string]any {
	var resources []map[string]any

	paths := []string{
		filepath.Join(root, actio.StandardFiles["index"]),
		filepath.Join(root, actio.StandardFiles["architecture"]),
		filepath.Join(root, actio.StandardFiles["interfaces"]),
		filepath.Join(root, actio.StandardFiles["rules"]),
		filepath.Join(root, actio.StandardFiles["tasks"]),
		filepath.Join(root, actio.StandardFiles["scripts_manifest"]),
	}

	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			rel, _ := filepath.Rel(root, p)
			resources = append(resources, map[string]any{
				"uri":         "actio://" + filepath.ToSlash(rel),
				"description": "Actio resource: " + rel,
			})
		}
	}

	return resources
}

func readResource(root, uri string) (string, error) {
	const prefix = "actio://"
	if uri == "" || len(uri) <= len(prefix) || uri[:len(prefix)] != prefix {
		return "", fmt.Errorf("unsupported URI (expected actio://): %s", uri)
	}
	rel := filepath.FromSlash(uri[len(prefix):])
	path := filepath.Join(root, rel)
	path = filepath.Clean(path)
	rootAbs, err := filepath.Abs(root)
	if err != nil {
		return "", fmt.Errorf("unable to resolve root: %w", err)
	}
	pathAbs, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("unable to resolve path: %w", err)
	}
	relToRoot, err := filepath.Rel(rootAbs, pathAbs)
	if err != nil || strings.HasPrefix(relToRoot, "..") {
		return "", fmt.Errorf("unsupported URI (path escapes root): %s", uri)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("unable to read resource: %w", err)
	}
	return string(data), nil
}

