package mcp

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

type fakePluginClient struct {
	resources []map[string]any
	content   map[string]string
}

func (f *fakePluginClient) ListResources(ctx context.Context) ([]map[string]any, error) {
	return f.resources, nil
}
func (f *fakePluginClient) ReadResource(ctx context.Context, uri string) (string, error) {
	return f.content[uri], nil
}
func (f *fakePluginClient) Close() error { return nil }

func TestLoadPluginConfigs_None(t *testing.T) {
	dir := t.TempDir()
	cfgs, err := loadPluginConfigs(dir)
	if err != nil {
		t.Fatalf("loadPluginConfigs: %v", err)
	}
	if len(cfgs) != 0 {
		t.Fatalf("expected 0 configs, got %d", len(cfgs))
	}
}

func TestLoadPluginConfigs_YAML_ExpandEnv(t *testing.T) {
	dir := t.TempDir()
	pluginsDir := filepath.Join(dir, "mcp", "plugins")
	if err := os.MkdirAll(pluginsDir, 0o755); err != nil {
		t.Fatal(err)
	}
	_ = os.Setenv("MCP_TEST_BIN", "echo")
	data := []byte(`
name: test
command: ${MCP_TEST_BIN}
args: ["hello", "${MCP_TEST_ARG}"]
env:
  FOO: ${MCP_TEST_FOO}
`)
	_ = os.Setenv("MCP_TEST_ARG", "world")
	_ = os.Setenv("MCP_TEST_FOO", "bar")
	if err := os.WriteFile(filepath.Join(pluginsDir, "test.yaml"), data, 0o644); err != nil {
		t.Fatal(err)
	}
	cfgs, err := loadPluginConfigs(dir)
	if err != nil {
		t.Fatalf("loadPluginConfigs: %v", err)
	}
	if len(cfgs) != 1 {
		t.Fatalf("expected 1 config, got %d", len(cfgs))
	}
	if cfgs[0].Command != "echo" {
		t.Fatalf("expected expanded command, got %q", cfgs[0].Command)
	}
	if got := cfgs[0].Args[1]; got != "world" {
		t.Fatalf("expected expanded arg, got %q", got)
	}
	if got := cfgs[0].Env["FOO"]; got != "bar" {
		t.Fatalf("expected expanded env, got %q", got)
	}
}

func TestPluginURI_RoundTrip(t *testing.T) {
	u := makePluginURI("p", "mcp://x/y?z=1")
	name, orig, ok := parsePluginURI(u)
	if !ok {
		t.Fatalf("expected ok")
	}
	if name != "p" || orig != "mcp://x/y?z=1" {
		t.Fatalf("round trip mismatch: %q %q", name, orig)
	}
}

func TestPluginManager_ListAllResources_RewriteURI(t *testing.T) {
	dir := t.TempDir()
	pluginsDir := filepath.Join(dir, "mcp", "plugins")
	if err := os.MkdirAll(pluginsDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(pluginsDir, "p.yaml"), []byte("name: p\ncommand: echo\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	pm := newPluginManager(dir, os.Stderr, func(ctx context.Context, cfg pluginConfig) (pluginClient, error) {
		return &fakePluginClient{
			resources: []map[string]any{
				{"uri": "mcp://resource/foo", "description": "Foo"},
			},
			content: map[string]string{"mcp://resource/foo": "ok"},
		}, nil
	})

	res, err := pm.listAllResources(context.Background())
	if err != nil {
		t.Fatalf("listAllResources: %v", err)
	}
	if len(res) != 1 {
		t.Fatalf("expected 1 resource, got %d", len(res))
	}
	uri, _ := res[0]["uri"].(string)
	name, orig, ok := parsePluginURI(uri)
	if !ok || name != "p" || orig != "mcp://resource/foo" {
		t.Fatalf("unexpected rewritten uri: %v", uri)
	}
}

