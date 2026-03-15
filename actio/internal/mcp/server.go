package mcp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

// ServeStdIO starts a minimal MCP-style JSON-RPC server over stdio.
// It intentionally supports a narrow tool surface focused on Actio context:
// - mcp.listResources
// - mcp.readResource
func ServeStdIO(ctx context.Context, root string, stderr io.Writer) error {
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

		resp := handleRequest(root, req)
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

func handleRequest(root string, req request) response {
	switch req.Method {
	case "mcp.listResources":
		return response{
			JsonRPC: "2.0",
			ID:      req.ID,
			Result: map[string]any{
				"resources": listResources(root),
			},
		}
	case "mcp.readResource":
		var params struct {
			URI string `json:"uri"`
		}
		if err := json.Unmarshal(req.Params, &params); err != nil {
			return errorResponse(req.ID, -32602, "invalid params")
		}
		content, err := readResource(root, params.URI)
		if err != nil {
			return errorResponse(req.ID, -32001, err.Error())
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
	rel := uri[len(prefix):]
	path := filepath.Join(root, filepath.FromSlash(rel))
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("unable to read resource: %w", err)
	}
	return string(data), nil
}

