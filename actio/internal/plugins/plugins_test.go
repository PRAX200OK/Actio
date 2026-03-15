package plugins

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunValidationPlugins_NoPluginsDir(t *testing.T) {
	dir := t.TempDir()
	// No actio/plugins
	issues, err := RunValidationPlugins(dir)
	if err != nil {
		t.Fatalf("RunValidationPlugins: %v", err)
	}
	if len(issues) != 0 {
		t.Errorf("expected no issues when no plugins dir: %v", issues)
	}
}

func TestRunValidationPlugins_EmptyPluginsDir(t *testing.T) {
	dir := t.TempDir()
	pluginsDir := filepath.Join(dir, "actio", "plugins")
	if err := os.MkdirAll(pluginsDir, 0o755); err != nil {
		t.Fatal(err)
	}
	issues, err := RunValidationPlugins(dir)
	if err != nil {
		t.Fatalf("RunValidationPlugins: %v", err)
	}
	if len(issues) != 0 {
		t.Errorf("expected no issues: %v", issues)
	}
}

func TestRunValidationPlugins_PluginReportsMissingFile(t *testing.T) {
	dir := t.TempDir()
	pluginsDir := filepath.Join(dir, "actio", "plugins")
	if err := os.MkdirAll(pluginsDir, 0o755); err != nil {
		t.Fatal(err)
	}
	yaml := `name: test-plugin
description: Requires a file
requiredFiles:
  - actio/plugins/required.txt
`
	if err := os.WriteFile(filepath.Join(pluginsDir, "test.yaml"), []byte(yaml), 0o644); err != nil {
		t.Fatal(err)
	}
	issues, err := RunValidationPlugins(dir)
	if err != nil {
		t.Fatalf("RunValidationPlugins: %v", err)
	}
	if len(issues) == 0 {
		t.Fatal("expected issue for missing required file")
	}
	var found bool
	for _, s := range issues {
		if strings.Contains(s, "plugin") && strings.Contains(s, "missing") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected plugin-related issue in: %v", issues)
	}
}

func TestRunValidationPlugins_PluginPassesWhenFilesExist(t *testing.T) {
	dir := t.TempDir()
	pluginsDir := filepath.Join(dir, "actio", "plugins")
	if err := os.MkdirAll(pluginsDir, 0o755); err != nil {
		t.Fatal(err)
	}
	requiredPath := filepath.Join(dir, "actio", "plugins", "required.txt")
	if err := os.WriteFile(requiredPath, []byte("ok"), 0o644); err != nil {
		t.Fatal(err)
	}
	yaml := `name: ok-plugin
requiredFiles:
  - actio/plugins/required.txt
`
	if err := os.WriteFile(filepath.Join(pluginsDir, "ok.yaml"), []byte(yaml), 0o644); err != nil {
		t.Fatal(err)
	}
	issues, err := RunValidationPlugins(dir)
	if err != nil {
		t.Fatalf("RunValidationPlugins: %v", err)
	}
	if len(issues) != 0 {
		t.Errorf("expected no issues when required file exists: %v", issues)
	}
}
