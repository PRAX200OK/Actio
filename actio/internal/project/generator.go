package project

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"actio/internal/actio"
	"actio/internal/templates"
)

// Preset defines the project structure to generate (minimal, standard, or full).
type Preset int

const (
	PresetMinimal Preset = iota
	PresetStandard
	PresetFull
)

func (p Preset) String() string {
	switch p {
	case PresetMinimal:
		return "minimal"
	case PresetStandard:
		return "standard"
	case PresetFull:
		return "full"
	default:
		return "unknown"
	}
}

// ParsePreset converts a string (minimal, standard, full) to Preset. Returns error if invalid.
func ParsePreset(s string) (Preset, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "minimal":
		return PresetMinimal, nil
	case "standard":
		return PresetStandard, nil
	case "full":
		return PresetFull, nil
	default:
		return PresetStandard, fmt.Errorf("invalid preset %q: must be minimal, standard, or full", s)
	}
}

// CreateNewProject generates a new project with Actio sidecar structure for the given preset.
func CreateNewProject(baseDir, name string, preset Preset) error {
	projectRoot := filepath.Join(baseDir, name)

	if _, err := os.Stat(projectRoot); err == nil {
		return fmt.Errorf("directory %s already exists", projectRoot)
	}

	if err := os.MkdirAll(filepath.Join(projectRoot, "src"), 0o755); err != nil {
		return fmt.Errorf("create src directory: %w", err)
	}

	if err := generateActSidecar(projectRoot, preset); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(projectRoot, actio.MainDoc), []byte(templates.ActMD()), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", actio.MainDoc, err)
	}

	return nil
}

// InitExistingRepo adds the Actio framework sidecar to an existing repository for the given preset.
func InitExistingRepo(root string, preset Preset) error {
	if err := generateActSidecar(root, preset); err != nil {
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

func generateActSidecar(root string, preset Preset) error {
	dirs := dirsForPreset(root, preset)
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0o755); err != nil {
			return fmt.Errorf("create directory %s: %w", d, err)
		}
	}

	files := filesForPreset(root, preset)
	for _, f := range files {
		if _, err := os.Stat(f.path); os.IsNotExist(err) {
			if err := os.WriteFile(f.path, []byte(f.content), 0o644); err != nil {
				return fmt.Errorf("write %s: %w", f.path, err)
			}
		}
	}

	return nil
}

func dirsForPreset(root string, preset Preset) []string {
	// Minimal: core only (no scripts). Standard: + scripts. Full: + actio/plugins, mcp/plugins.
	dirs := []string{
		filepath.Join(root, actio.ActioPath("architecture")),
		filepath.Join(root, actio.ActioPath("interfaces")),
		filepath.Join(root, actio.ActioPath("patterns")),
		filepath.Join(root, actio.ActioPath("rules")),
		filepath.Join(root, actio.ActioPath("tasks")),
	}
	if preset == PresetMinimal {
		return dirs
	}
	dirs = append(dirs, filepath.Join(root, actio.ScriptsDir))
	if preset == PresetFull {
		dirs = append(dirs, filepath.Join(root, actio.ActioPath("plugins")))
		dirs = append(dirs, filepath.Join(root, "mcp", "plugins"))
	}
	return dirs
}

type fileEntry struct {
	path    string
	content string
}

func filesForPreset(root string, preset Preset) []fileEntry {
	files := []fileEntry{
		{filepath.Join(root, actio.StandardFiles["index"]), templates.IndexYAML},
		{filepath.Join(root, actio.StandardFiles["architecture"]), templates.ArchitectureSystemMD},
		{filepath.Join(root, actio.StandardFiles["interfaces"]), templates.InterfacesContractsYAML},
		{filepath.Join(root, actio.StandardFiles["patterns"]), templates.PatternsExampleMD},
		{filepath.Join(root, actio.StandardFiles["rules"]), templates.RulesCodingMD()},
		{filepath.Join(root, actio.StandardFiles["tasks"]), templates.TasksExampleMD()},
	}
	if preset != PresetMinimal {
		files = append(files,
			fileEntry{filepath.Join(root, actio.StandardFiles["scripts_manifest"]), templates.ScriptsManifestYAML},
			fileEntry{filepath.Join(root, actio.ScriptsDir, "example.py"), templates.ExampleScriptPy},
		)
	}
	if preset == PresetFull {
		files = append(files,
			fileEntry{filepath.Join(root, actio.ActioPath("plugins", "README.md")), templates.ActioPluginsReadmeMD()},
			fileEntry{filepath.Join(root, "mcp", "plugins", "README.md"), templates.MCPPluginsReadmeMD()},
		)
	}
	return files
}

