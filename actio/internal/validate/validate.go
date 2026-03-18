package validate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"actio/internal/actio"
	"actio/internal/plugins"

	"gopkg.in/yaml.v3"
)

// Index models the Actio router.yaml schema for validation.
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
			issues = append(issues, fmt.Sprintf("missing %s", actio.StandardFiles["index"]))
		} else {
			return nil, fmt.Errorf("stat %s: %w", actio.IndexFile, err)
		}
	} else {
		data, err := os.ReadFile(indexPath)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", actio.IndexFile, err)
		}
		if err := yaml.Unmarshal(data, &idx); err != nil {
			issues = append(issues, fmt.Sprintf("%s is not valid YAML", actio.IndexFile))
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
		filepath.Join(root, actio.StandardFiles["interfaces"]),
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

// pathUnderActRoot reports whether cleanPath is under actRoot (no path traversal out of actio/).
func pathUnderActRoot(actRoot, cleanPath string) bool {
	rel, err := filepath.Rel(actRoot, cleanPath)
	if err != nil {
		return false
	}
	return rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator))
}

func validateIndexSchema(root, actRoot string, idx Index) []string {
	var issues []string

	if idx.Version == 0 {
		issues = append(issues, fmt.Sprintf("%s: version must be set and > 0", actio.IndexFile))
	}
	if idx.Project.Name == "" {
		issues = append(issues, fmt.Sprintf("%s: project.name must be set", actio.IndexFile))
	}

	// Validate domains
	for name, d := range idx.Domains {
		if d.Architecture == "" {
			issues = append(issues, fmt.Sprintf("%s: domains.%s.architecture must be set", actio.IndexFile, name))
		} else {
			path := filepath.Clean(filepath.Join(actRoot, d.Architecture))
			if !pathUnderActRoot(actRoot, path) {
				issues = append(issues, fmt.Sprintf("%s: domains.%s.architecture path escapes actio/", actio.IndexFile, name))
			} else if _, err := os.Stat(path); err != nil {
				rel, _ := filepath.Rel(root, path)
				issues = append(issues, fmt.Sprintf("%s: domains.%s.architecture references missing file: %s", actio.IndexFile, name, rel))
			}
		}

		if d.Interfaces == "" {
			issues = append(issues, fmt.Sprintf("%s: domains.%s.interfaces must be set", actio.IndexFile, name))
		} else {
			path := filepath.Clean(filepath.Join(actRoot, d.Interfaces))
			if !pathUnderActRoot(actRoot, path) {
				issues = append(issues, fmt.Sprintf("%s: domains.%s.interfaces path escapes actio/", actio.IndexFile, name))
			} else if _, err := os.Stat(path); err != nil {
				rel, _ := filepath.Rel(root, path)
				issues = append(issues, fmt.Sprintf("%s: domains.%s.interfaces references missing file: %s", actio.IndexFile, name, rel))
			}
		}

		for _, p := range d.Patterns {
			path := filepath.Clean(filepath.Join(actRoot, p))
			if !pathUnderActRoot(actRoot, path) {
				issues = append(issues, fmt.Sprintf("%s: domains.%s.patterns path escapes actio/", actio.IndexFile, name))
			} else if _, err := os.Stat(path); err != nil {
				rel, _ := filepath.Rel(root, path)
				issues = append(issues, fmt.Sprintf("%s: domains.%s.patterns references missing file: %s", actio.IndexFile, name, rel))
			}
		}
	}

	// Validate rules
	if idx.Rules.Coding == "" {
		issues = append(issues, fmt.Sprintf("%s: rules.coding must be set", actio.IndexFile))
	} else {
		path := filepath.Clean(filepath.Join(actRoot, idx.Rules.Coding))
		if !pathUnderActRoot(actRoot, path) {
			issues = append(issues, fmt.Sprintf("%s: rules.coding path escapes actio/", actio.IndexFile))
		} else if _, err := os.Stat(path); err != nil {
			rel, _ := filepath.Rel(root, path)
			issues = append(issues, fmt.Sprintf("%s: rules.coding references missing file: %s", actio.IndexFile, rel))
		}
	}

	// Validate tasks
	for taskName, t := range idx.Tasks {
		if t.Domain == "" {
			issues = append(issues, fmt.Sprintf("%s: tasks.%s.domain must be set", actio.IndexFile, taskName))
		} else {
			if _, ok := idx.Domains[t.Domain]; !ok {
				issues = append(issues, fmt.Sprintf("%s: tasks.%s.domain references unknown domain: %s", actio.IndexFile, taskName, t.Domain))
			}
		}
		if t.Guide == "" {
			issues = append(issues, fmt.Sprintf("%s: tasks.%s.guide must be set", actio.IndexFile, taskName))
		} else {
			path := filepath.Clean(filepath.Join(actRoot, t.Guide))
			if !pathUnderActRoot(actRoot, path) {
				issues = append(issues, fmt.Sprintf("%s: tasks.%s.guide path escapes actio/", actio.IndexFile, taskName))
			} else if _, err := os.Stat(path); err != nil {
				rel, _ := filepath.Rel(root, path)
				issues = append(issues, fmt.Sprintf("%s: tasks.%s.guide references missing file: %s", actio.IndexFile, taskName, rel))
			}
		}
	}

	return issues
}


