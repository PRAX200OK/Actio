package project

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateNewProject_Success(t *testing.T) {
	base := t.TempDir()
	name := "myapp"
	err := CreateNewProject(base, name)
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
	for _, d := range []string{"architecture", "interfaces", "patterns", "rules", "tasks"} {
		p := filepath.Join(actRoot, d)
		if fi, err := os.Stat(p); err != nil || !fi.IsDir() {
			t.Errorf("expected directory %s: err=%v isDir=%v", p, err, fi != nil && fi.IsDir())
		}
	}
	for _, f := range []string{"actio/router.yaml", "actio/architecture/system.md", "actio/rules/rules.md", "actio/tasks/task.md"} {
		p := filepath.Join(root, f)
		if _, err := os.Stat(p); err != nil {
			t.Errorf("expected file %s: %v", p, err)
		}
	}
}

func TestCreateNewProject_AlreadyExists(t *testing.T) {
	base := t.TempDir()
	name := "existing"
	root := filepath.Join(base, name)
	if err := os.MkdirAll(root, 0o755); err != nil {
		t.Fatal(err)
	}
	err := CreateNewProject(base, name)
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

func TestInitExistingRepo_Success(t *testing.T) {
	dir := t.TempDir()
	err := InitExistingRepo(dir)
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

func TestInitExistingRepo_DoesNotOverwriteACTIOMd(t *testing.T) {
	dir := t.TempDir()
	actPath := filepath.Join(dir, "ENTRYPOINT.yaml")
	custom := "# Custom ENTRYPOINT"
	if err := os.WriteFile(actPath, []byte(custom), 0o644); err != nil {
		t.Fatal(err)
	}
	err := InitExistingRepo(dir)
	if err != nil {
		t.Fatalf("InitExistingRepo: %v", err)
	}
	data, err := os.ReadFile(actPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != custom {
		t.Errorf("ACTIO.md was overwritten: got %q", string(data))
	}
}
