package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestListResources_EmptyRoot(t *testing.T) {
	dir := t.TempDir()
	res := listResources(dir)
	if len(res) != 0 {
		t.Errorf("expected no resources in empty root, got %d", len(res))
	}
}

func TestListResources_WithActFiles(t *testing.T) {
	dir := t.TempDir()
	actRoot := filepath.Join(dir, "actio")
	for _, d := range []string{"architecture", "interfaces", "rules", "tasks"} {
		_ = os.MkdirAll(filepath.Join(actRoot, d), 0o755)
	}
	for _, f := range []string{"router.yaml", "architecture/system.md", "interfaces/contracts.yaml", "rules/rules.md", "tasks/task.md"} {
		_ = os.WriteFile(filepath.Join(actRoot, f), []byte("x"), 0o644)
	}

	res := listResources(dir)
	if len(res) < 5 {
		t.Errorf("expected at least 5 resources, got %d", len(res))
	}
	const prefix = "actio://"
	for _, r := range res {
		uri, _ := r["uri"].(string)
		if uri == "" || len(uri) < len(prefix) || uri[:len(prefix)] != prefix {
			t.Errorf("resource missing or invalid uri: %v", r)
		}
	}
}

func TestReadResource_Success(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "actio", "router.yaml")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	content := "version: 1"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	uri := "actio://actio/router.yaml"
	got, err := readResource(dir, uri)
	if err != nil {
		t.Fatalf("readResource: %v", err)
	}
	if got != content {
		t.Errorf("got %q want %q", got, content)
	}
}

func TestReadResource_InvalidURI(t *testing.T) {
	dir := t.TempDir()
	_, err := readResource(dir, "http://example.com/act/index.yaml")
	if err == nil {
		t.Fatal("expected error for non-act URI")
	}
	_, err = readResource(dir, "")
	if err == nil {
		t.Fatal("expected error for empty URI")
	}
}

func TestReadResource_NotFound(t *testing.T) {
	dir := t.TempDir()
	_, err := readResource(dir, "actio://actio/nonexistent.yaml")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestHandleRequest_ListResources(t *testing.T) {
	dir := t.TempDir()
	actRoot := filepath.Join(dir, "actio")
	_ = os.MkdirAll(filepath.Join(actRoot, "architecture"), 0o755)
	_ = os.WriteFile(filepath.Join(actRoot, "router.yaml"), []byte("x"), 0o644)

	req := request{JsonRPC: "2.0", ID: json.RawMessage(`1`), Method: "mcp.listResources", Params: json.RawMessage(`{}`)}
	s := &server{root: dir, pluginMgr: newPluginManager(dir, io.Discard, func(ctx context.Context, cfg pluginConfig) (pluginClient, error) {
		return nil, fmt.Errorf("no plugins")
	})}
	resp := s.handleRequest(context.Background(), req)
	if resp.Error != nil {
		t.Fatalf("unexpected error: %v", resp.Error)
	}
	m, _ := resp.Result.(map[string]any)
	resources, _ := m["resources"].([]map[string]any)
	if len(resources) == 0 {
		t.Error("expected at least one resource")
	}
}

func TestHandleRequest_ReadResource(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "actio", "router.yaml")
	_ = os.MkdirAll(filepath.Dir(path), 0o755)
	_ = os.WriteFile(path, []byte("hello"), 0o644)

	params, _ := json.Marshal(map[string]string{"uri": "actio://actio/router.yaml"})
	req := request{JsonRPC: "2.0", ID: json.RawMessage(`2`), Method: "mcp.readResource", Params: params}
	s := &server{root: dir, pluginMgr: newPluginManager(dir, io.Discard, func(ctx context.Context, cfg pluginConfig) (pluginClient, error) {
		return nil, fmt.Errorf("no plugins")
	})}
	resp := s.handleRequest(context.Background(), req)
	if resp.Error != nil {
		t.Fatalf("unexpected error: %v", resp.Error)
	}
	m, _ := resp.Result.(map[string]any)
	if m["content"] != "hello" {
		t.Errorf("content: got %v", m["content"])
	}
}

func TestHandleRequest_MethodNotFound(t *testing.T) {
	req := request{JsonRPC: "2.0", ID: json.RawMessage(`3`), Method: "unknown.method"}
	s := &server{root: t.TempDir(), pluginMgr: newPluginManager(t.TempDir(), io.Discard, func(ctx context.Context, cfg pluginConfig) (pluginClient, error) {
		return nil, fmt.Errorf("no plugins")
	})}
	resp := s.handleRequest(context.Background(), req)
	if resp.Error == nil {
		t.Fatal("expected error for unknown method")
	}
	if resp.Error.Code != -32601 {
		t.Errorf("expected -32601, got %d", resp.Error.Code)
	}
}
