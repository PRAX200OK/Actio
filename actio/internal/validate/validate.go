package validate

import (
	"fmt"
	"os"
	"path/filepath"

	"actio/internal/plugins"
	"actio/internal/actio"

	"gopkg.in/yaml.v3"
)

// Index models the Actio index.yaml schema for validation.
type Index struct {
	Version int `yaml:"version"`

	Project struct {
		Name string `yaml:"name"`
	} `yaml:"project"`

	Domains map[string]struct {
		Architecture string   `yaml:"architecture"`
		Interfaces   string   `yaml:"interfaces"`
		Patterns     []string `yaml:"patterns"`
	} `yaml:"domains"`

	Rules struct {
		Coding string `yaml:"coding"`
	} `yaml:"rules"`

	Tasks map[string]struct {
		Domain string `yaml:"domain"`
		Guide  string `yaml:"guide"`
	} `yaml:"tasks"`
}

// Validate runs core Actio validation rules (structure + schema + referential integrity)
// and plugin-defined checks. An empty slice means validation passed.
func Validate(root string) ([]string, error) {
	var issues []string

	actRoot := filepath.Join(root, actio.DirName)
	indexPath := filepath.Join(root, actio.StandardFiles["index"])

	var idx Index

	if _, err := os.Stat(indexPath); err != nil {
		if os.IsNotExist(err) {
			issues = append(issues, "missing actio/index.yaml")
		} else {
			return nil, fmt.Errorf("stat index.yaml: %w", err)
		}
	} else {
		data, err := os.ReadFile(indexPath)
		if err != nil {
			return nil, fmt.Errorf("read index.yaml: %w", err)
		}
		if err := yaml.Unmarshal(data, &idx); err != nil {
			issues = append(issues, "index.yaml is not valid YAML")
		} else {
			issues = append(issues, validateIndexSchema(root, actRoot, idx)...)
		}
	}

	requiredDirs := []string{
		filepath.Join(actRoot, "architecture"),
		filepath.Join(actRoot, "interfaces"),
		filepath.Join(actRoot, "rules"),
		filepath.Join(actRoot, "tasks"),
	}

	for _, d := range requiredDirs {
		if fi, err := os.Stat(d); err != nil || !fi.IsDir() {
			rel, _ := filepath.Rel(root, d)
			issues = append(issues, fmt.Sprintf("missing required directory: %s", rel))
		}
	}

	// Basic module files
	requiredFiles := []string{
		filepath.Join(root, actio.StandardFiles["architecture"]),
		filepath.Join(root, actio.StandardFiles["rules"]),
		filepath.Join(root, actio.StandardFiles["tasks"]),
	}

	for _, f := range requiredFiles {
		if _, err := os.Stat(f); err != nil {
			rel, _ := filepath.Rel(root, f)
			issues = append(issues, fmt.Sprintf("missing required file: %s", rel))
		}
	}

	// Run plugin-based validations
	pluginIssues, err := plugins.RunValidationPlugins(root)
	if err != nil {
		return nil, err
	}
	issues = append(issues, pluginIssues...)

	return issues, nil
}

func validateIndexSchema(root, actRoot string, idx Index) []string {
	var issues []string

	if idx.Version == 0 {
		issues = append(issues, "index.yaml: version must be set and > 0")
	}
	if idx.Project.Name == "" {
		issues = append(issues, "index.yaml: project.name must be set")
	}

	// Validate domains
	for name, d := range idx.Domains {
		if d.Architecture == "" {
			issues = append(issues, fmt.Sprintf("index.yaml: domains.%s.architecture must be set", name))
		} else {
			path := filepath.Join(actRoot, d.Architecture)
			if _, err := os.Stat(path); err != nil {
				rel, _ := filepath.Rel(root, path)
				issues = append(issues, fmt.Sprintf("index.yaml: domains.%s.architecture references missing file: %s", name, rel))
			}
		}

		if d.Interfaces == "" {
			issues = append(issues, fmt.Sprintf("index.yaml: domains.%s.interfaces must be set", name))
		} else {
			path := filepath.Join(actRoot, d.Interfaces)
			if _, err := os.Stat(path); err != nil {
				rel, _ := filepath.Rel(root, path)
				issues = append(issues, fmt.Sprintf("index.yaml: domains.%s.interfaces references missing file: %s", name, rel))
			}
		}

		for _, p := range d.Patterns {
			path := filepath.Join(actRoot, p)
			if _, err := os.Stat(path); err != nil {
				rel, _ := filepath.Rel(root, path)
				issues = append(issues, fmt.Sprintf("index.yaml: domains.%s.patterns references missing file: %s", name, rel))
			}
		}
	}

	// Validate rules
	if idx.Rules.Coding == "" {
		issues = append(issues, "index.yaml: rules.coding must be set")
	} else {
		path := filepath.Join(actRoot, idx.Rules.Coding)
		if _, err := os.Stat(path); err != nil {
			rel, _ := filepath.Rel(root, path)
			issues = append(issues, fmt.Sprintf("index.yaml: rules.coding references missing file: %s", rel))
		}
	}

	// Validate tasks
	for taskName, t := range idx.Tasks {
		if t.Domain == "" {
			issues = append(issues, fmt.Sprintf("index.yaml: tasks.%s.domain must be set", taskName))
		} else {
			if _, ok := idx.Domains[t.Domain]; !ok {
				issues = append(issues, fmt.Sprintf("index.yaml: tasks.%s.domain references unknown domain: %s", taskName, t.Domain))
			}
		}
		if t.Guide == "" {
			issues = append(issues, fmt.Sprintf("index.yaml: tasks.%s.guide must be set", taskName))
		} else {
			path := filepath.Join(actRoot, t.Guide)
			if _, err := os.Stat(path); err != nil {
				rel, _ := filepath.Rel(root, path)
				issues = append(issues, fmt.Sprintf("index.yaml: tasks.%s.guide references missing file: %s", taskName, rel))
			}
		}
	}

	return issues
}


