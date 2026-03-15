package validate

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidate_MissingActDir(t *testing.T) {
	dir := t.TempDir()
	// No actio/ at all
	issues, err := Validate(dir)
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	if len(issues) == 0 {
		t.Fatal("expected issues when actio/ is missing")
	}
	var found bool
	for _, s := range issues {
		if s == "missing actio/index.yaml" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected 'missing actio/index.yaml' in issues: %v", issues)
	}
}

func TestValidate_InvalidYAML(t *testing.T) {
	dir := t.TempDir()
	actRoot := filepath.Join(dir, "actio")
	if err := os.MkdirAll(actRoot, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(actRoot, "index.yaml"), []byte("not: valid: yaml: ["), 0o644); err != nil {
		t.Fatal(err)
	}
	// Create required dirs/files so we don't get those issues
	for _, d := range []string{"architecture", "interfaces", "rules", "tasks"} {
		_ = os.MkdirAll(filepath.Join(actRoot, d), 0o755)
	}
	for _, f := range []string{"architecture/system.md", "rules/rules.md", "tasks/task.md"} {
		_ = os.WriteFile(filepath.Join(actRoot, f), []byte("x"), 0o644)
	}

	issues, err := Validate(dir)
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	var found bool
	for _, s := range issues {
		if s == "index.yaml is not valid YAML" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected invalid YAML issue in: %v", issues)
	}
}

func TestValidate_ValidProject(t *testing.T) {
	dir := t.TempDir()
	actRoot := filepath.Join(dir, "actio")
	if err := os.MkdirAll(filepath.Join(actRoot, "architecture"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(actRoot, "interfaces"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(actRoot, "rules"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(actRoot, "tasks"), 0o755); err != nil {
		t.Fatal(err)
	}
	indexYAML := `version: 1
project:
  name: test-project
domains:
  connectors:
    architecture: architecture/system.md
    interfaces: interfaces/contracts.yaml
    patterns: []
rules:
  coding: rules/rules.md
tasks:
  example_task:
    domain: connectors
    guide: tasks/task.md
`
	if err := os.WriteFile(filepath.Join(actRoot, "index.yaml"), []byte(indexYAML), 0o644); err != nil {
		t.Fatal(err)
	}
	for _, f := range []string{"architecture/system.md", "interfaces/contracts.yaml", "rules/rules.md", "tasks/task.md"} {
		if err := os.WriteFile(filepath.Join(actRoot, f), []byte("content"), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	issues, err := Validate(dir)
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	if len(issues) != 0 {
		t.Errorf("expected no issues for valid project, got: %v", issues)
	}
}

func TestValidate_SchemaErrors(t *testing.T) {
	dir := t.TempDir()
	actRoot := filepath.Join(dir, "actio")
	for _, d := range []string{"architecture", "interfaces", "rules", "tasks"} {
		if err := os.MkdirAll(filepath.Join(actRoot, d), 0o755); err != nil {
			t.Fatal(err)
		}
	}

	// Missing version and project.name; task references unknown domain
	indexYAML := `version: 0
project:
  name: ""
domains:
  connectors:
    architecture: architecture/system.md
    interfaces: interfaces/contracts.yaml
    patterns: []
rules:
  coding: rules/rules.md
tasks:
  bad_task:
    domain: nonexistent
    guide: tasks/task.md
`
	if err := os.WriteFile(filepath.Join(actRoot, "index.yaml"), []byte(indexYAML), 0o644); err != nil {
		t.Fatal(err)
	}
	for _, f := range []string{"architecture/system.md", "interfaces/contracts.yaml", "rules/rules.md", "tasks/task.md"} {
		if err := os.WriteFile(filepath.Join(actRoot, f), []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	issues, err := Validate(dir)
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	wantContains := []string{"version must be set", "project.name must be set", "unknown domain"}
	for _, want := range wantContains {
		var found bool
		for _, s := range issues {
			if containsSub(s, want) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected an issue containing %q in: %v", want, issues)
		}
	}
}

func containsSub(a, sub string) bool {
	for i := 0; i <= len(a)-len(sub); i++ {
		if a[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

func TestValidateIndexSchema_ReferentialIntegrity(t *testing.T) {
	dir := t.TempDir()
	actRoot := filepath.Join(dir, "act")
	if err := os.MkdirAll(filepath.Join(actRoot, "architecture"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(actRoot, "architecture/system.md"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	// interfaces file missing
	idx := Index{}
	idx.Version = 1
	idx.Project.Name = "p"
	idx.Domains = map[string]struct {
		Architecture string   `yaml:"architecture"`
		Interfaces   string   `yaml:"interfaces"`
		Patterns     []string `yaml:"patterns"`
	}{
		"connectors": {
			Architecture: "architecture/system.md",
			Interfaces:   "interfaces/nonexistent.yaml",
			Patterns:     nil,
		},
	}
	idx.Rules.Coding = "rules/rules.md"
	idx.Tasks = map[string]struct {
		Domain string `yaml:"domain"`
		Guide  string `yaml:"guide"`
	}{
		"t1": {Domain: "connectors", Guide: "tasks/t1.md"},
	}
	issues := validateIndexSchema(dir, actRoot, idx)
	var refErr bool
	for _, s := range issues {
		if containsSub(s, "references missing file") {
			refErr = true
			break
		}
	}
	if !refErr {
		t.Errorf("expected referential error for missing interfaces/task file, got: %v", issues)
	}
}
