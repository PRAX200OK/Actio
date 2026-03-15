package plugins

import (
	"fmt"
	"os"
	"path/filepath"

	"actio/internal/actio"
	"gopkg.in/yaml.v3"
)

// ValidationPlugin describes a simple, YAML-defined plugin that can add extra checks.
type ValidationPlugin struct {
	Name          string   `yaml:"name"`
	Description   string   `yaml:"description"`
	RequiredFiles []string `yaml:"requiredFiles"`
}

// RunValidationPlugins discovers and runs validation plugins under actio/plugins.
// It returns a list of human-readable issues.
func RunValidationPlugins(root string) ([]string, error) {
	var issues []string

	pluginsDir := filepath.Join(root, actio.ActioPath("plugins"))

	if _, err := os.Stat(pluginsDir); err != nil {
		// No plugins directory is not an error; just means no plugins to run.
		return issues, nil
	}

	entries, err := os.ReadDir(pluginsDir)
	if err != nil {
		return nil, fmt.Errorf("read plugins directory: %w", err)
	}

	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		if filepath.Ext(e.Name()) != ".yaml" && filepath.Ext(e.Name()) != ".yml" {
			continue
		}
		path := filepath.Join(pluginsDir, e.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read plugin %s: %w", e.Name(), err)
		}
		var p ValidationPlugin
		if err := yaml.Unmarshal(data, &p); err != nil {
			return nil, fmt.Errorf("parse plugin %s: %w", e.Name(), err)
		}
		if p.Name == "" {
			p.Name = e.Name()
		}

		for _, rf := range p.RequiredFiles {
			full := filepath.Join(root, rf)
			if _, err := os.Stat(full); err != nil {
				rel, _ := filepath.Rel(root, full)
				issues = append(issues, fmt.Sprintf("plugin %q: missing required file: %s", p.Name, rel))
			}
		}
	}

	return issues, nil
}

