package project

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateNewProject_Success(t *testing.T) {
	base := t.TempDir()
	name := "myapp"
	err := CreateNewProject(base, name, PresetStandard)
	if err != nil {
		t.Fatalf("CreateNewProject: %v", err)
	}
	root := filepath.Join(base, name)
	if _, err := os.Stat(root); os.IsNotExist(err) {
		t.Fatal("project root directory was not created")
	}
	if _, err := os.Stat(filepath.Join(root, "src")); os.IsNotExist(err) {
		t.Fatal("src directory was not created")
	}
	if _, err := os.Stat(filepath.Join(root, "ENTRYPOINT.yaml")); os.IsNotExist(err) {
		t.Fatal("ENTRYPOINT.yaml was not created")
	}
	actRoot := filepath.Join(root, "actio")
	for _, d := range []string{"architecture", "interfaces", "patterns", "rules", "tasks", "scripts"} {
		p := filepath.Join(actRoot, d)
		if fi, err := os.Stat(p); err != nil || !fi.IsDir() {
			t.Errorf("expected directory %s: err=%v isDir=%v", p, err, fi != nil && fi.IsDir())
		}
	}
	for _, f := range []string{"actio/router.yaml", "actio/architecture/system.md", "actio/rules/rules.md", "actio/tasks/task.md", "actio/scripts/manifest.yaml", "actio/scripts/example.py"} {
		p := filepath.Join(root, f)
		if _, err := os.Stat(p); err != nil {
			t.Errorf("expected file %s: %v", p, err)
		}
	}
}

func TestCreateNewProject_MinimalPreset(t *testing.T) {
	base := t.TempDir()
	name := "minimalapp"
	err := CreateNewProject(base, name, PresetMinimal)
	if err != nil {
		t.Fatalf("CreateNewProject: %v", err)
	}
	root := filepath.Join(base, name)
	actRoot := filepath.Join(root, "actio")
	for _, d := range []string{"architecture", "interfaces", "patterns", "rules", "tasks"} {
		p := filepath.Join(actRoot, d)
		if fi, err := os.Stat(p); err != nil || !fi.IsDir() {
			t.Errorf("expected directory %s: err=%v isDir=%v", p, err, fi != nil && fi.IsDir())
		}
	}
	scriptsDir := filepath.Join(actRoot, "scripts")
	if fi, err := os.Stat(scriptsDir); err == nil && fi.IsDir() {
		t.Error("minimal preset should not create actio/scripts/")
	}
	for _, f := range []string{"actio/router.yaml", "actio/architecture/system.md", "actio/interfaces/contracts.yaml", "actio/rules/rules.md", "actio/tasks/task.md"} {
		p := filepath.Join(root, f)
		if _, err := os.Stat(p); err != nil {
			t.Errorf("expected file %s: %v", p, err)
		}
	}
}

func TestCreateNewProject_FullPreset(t *testing.T) {
	base := t.TempDir()
	name := "fullapp"
	err := CreateNewProject(base, name, PresetFull)
	if err != nil {
		t.Fatalf("CreateNewProject: %v", err)
	}
	root := filepath.Join(base, name)
	actRoot := filepath.Join(root, "actio")
	if fi, err := os.Stat(filepath.Join(actRoot, "plugins")); err != nil || !fi.IsDir() {
		t.Errorf("full preset should create actio/plugins/: err=%v isDir=%v", err, fi != nil && fi.IsDir())
	}
	if fi, err := os.Stat(filepath.Join(root, "mcp", "plugins")); err != nil || !fi.IsDir() {
		t.Errorf("full preset should create mcp/plugins/: err=%v isDir=%v", err, fi != nil && fi.IsDir())
	}
	if _, err := os.Stat(filepath.Join(actRoot, "plugins", "README.md")); err != nil {
		t.Errorf("full preset should create actio/plugins/README.md: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "mcp", "plugins", "README.md")); err != nil {
		t.Errorf("full preset should create mcp/plugins/README.md: %v", err)
	}
}

func TestCreateNewProject_AlreadyExists(t *testing.T) {
	base := t.TempDir()
	name := "existing"
	root := filepath.Join(base, name)
	if err := os.MkdirAll(root, 0o755); err != nil {
		t.Fatal(err)
	}
	err := CreateNewProject(base, name, PresetStandard)
	if err == nil {
		t.Fatal("expected error when project directory already exists")
	}
	if !contains(err.Error(), "already exists") {
		t.Errorf("error should mention already exists: %v", err)
	}
}

func contains(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

func TestParsePreset(t *testing.T) {
	tests := []struct {
		in   string
		want Preset
		ok   bool
	}{
		{"minimal", PresetMinimal, true},
		{"standard", PresetStandard, true},
		{"full", PresetFull, true},
		{"  full  ", PresetFull, true},
		{"MINIMAL", PresetMinimal, true},
		{"", PresetStandard, false},
		{"invalid", PresetStandard, false},
	}
	for _, tt := range tests {
		got, err := ParsePreset(tt.in)
		if tt.ok {
			if err != nil {
				t.Errorf("ParsePreset(%q): %v", tt.in, err)
			}
			if got != tt.want {
				t.Errorf("ParsePreset(%q) = %v want %v", tt.in, got, tt.want)
			}
		} else {
			if err == nil {
				t.Errorf("ParsePreset(%q) expected error", tt.in)
			}
		}
	}
}

func TestInitExistingRepo_Success(t *testing.T) {
	dir := t.TempDir()
	err := InitExistingRepo(dir, PresetStandard)
	if err != nil {
		t.Fatalf("InitExistingRepo: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "ENTRYPOINT.yaml")); os.IsNotExist(err) {
		t.Fatal("ENTRYPOINT.yaml was not created")
	}
	if _, err := os.Stat(filepath.Join(dir, "actio", "router.yaml")); os.IsNotExist(err) {
		t.Fatal("actio/router.yaml was not created")
	}
}

func TestInitExistingRepo_MinimalPreset(t *testing.T) {
	dir := t.TempDir()
	err := InitExistingRepo(dir, PresetMinimal)
	if err != nil {
		t.Fatalf("InitExistingRepo: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "actio", "scripts")); err == nil {
		t.Error("minimal preset should not create actio/scripts/")
	}
	if _, err := os.Stat(filepath.Join(dir, "actio", "router.yaml")); err != nil {
		t.Errorf("actio/router.yaml should exist: %v", err)
	}
}

func TestInitExistingRepo_FullPreset(t *testing.T) {
	dir := t.TempDir()
	err := InitExistingRepo(dir, PresetFull)
	if err != nil {
		t.Fatalf("InitExistingRepo: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "actio", "plugins", "README.md")); err != nil {
		t.Errorf("full preset should create actio/plugins/README.md: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "mcp", "plugins", "README.md")); err != nil {
		t.Errorf("full preset should create mcp/plugins/README.md: %v", err)
	}
}

func TestInitExistingRepo_DoesNotOverwriteENTRYPOINT(t *testing.T) {
	dir := t.TempDir()
	actPath := filepath.Join(dir, "ENTRYPOINT.yaml")
	custom := "# Custom ENTRYPOINT"
	if err := os.WriteFile(actPath, []byte(custom), 0o644); err != nil {
		t.Fatal(err)
	}
	err := InitExistingRepo(dir, PresetStandard)
	if err != nil {
		t.Fatalf("InitExistingRepo: %v", err)
	}
	data, err := os.ReadFile(actPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != custom {
		t.Errorf("ENTRYPOINT.yaml was overwritten: got %q", string(data))
	}
}
