package project

import (
	"fmt"
	"os"
	"path/filepath"

	"actio/internal/templates"
	"actio/internal/actio"
)

// CreateNewProject generates a new project with ACT sidecar structure.
func CreateNewProject(baseDir, name string) error {
	projectRoot := filepath.Join(baseDir, name)

	if _, err := os.Stat(projectRoot); err == nil {
		return fmt.Errorf("directory %s already exists", projectRoot)
	}

	if err := os.MkdirAll(filepath.Join(projectRoot, "src"), 0o755); err != nil {
		return fmt.Errorf("create src directory: %w", err)
	}

	if err := generateActSidecar(projectRoot); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(projectRoot, actio.MainDoc), []byte(templates.ActMD()), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", actio.MainDoc, err)
	}

	return nil
}

// InitExistingRepo adds the ACTIO framework sidecar to an existing repository.
func InitExistingRepo(root string) error {
	if err := generateActSidecar(root); err != nil {
		return err
	}

	actPath := filepath.Join(root, actio.MainDoc)
	if _, err := os.Stat(actPath); os.IsNotExist(err) {
		if err := os.WriteFile(actPath, []byte(templates.ActMD()), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", actio.MainDoc, err)
		}
	}

	return nil
}

func generateActSidecar(root string) error {
	// Use ActioPath to build expected directory structure to avoid hard-coded paths.
	dirs := []string{
		filepath.Join(root, actio.ActioPath("architecture")),
		filepath.Join(root, actio.ActioPath("interfaces")),
		filepath.Join(root, actio.ActioPath("patterns")),
		filepath.Join(root, actio.ActioPath("rules")),
		filepath.Join(root, actio.ActioPath("tasks")),
	}

	for _, d := range dirs {
		if err := os.MkdirAll(d, 0o755); err != nil {
			return fmt.Errorf("create directory %s: %w", d, err)
		}
	}

	files := []struct {
		path    string
		content string
	}{
		{filepath.Join(root, actio.StandardFiles["index"]), templates.IndexYAML},
		{filepath.Join(root, actio.StandardFiles["architecture"]), templates.ArchitectureSystemMD},
		{filepath.Join(root, actio.StandardFiles["interfaces"]), templates.InterfacesContractsYAML},
		{filepath.Join(root, actio.StandardFiles["patterns"]), templates.PatternsExampleMD},
		{filepath.Join(root, actio.StandardFiles["rules"]), templates.RulesCodingMD()},
		{filepath.Join(root, actio.StandardFiles["tasks"]), templates.TasksExampleMD()},
	}

	for _, f := range files {
		if _, err := os.Stat(f.path); os.IsNotExist(err) {
			if err := os.WriteFile(f.path, []byte(f.content), 0o644); err != nil {
				return fmt.Errorf("write %s: %w", f.path, err)
			}
		}
	}

	return nil
}

