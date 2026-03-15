package actio

import "path/filepath"

// ActioPath joins the Actio root directory with the provided path elements.
// It is intended to be used anywhere you need a consistent path under the actio/ tree.
func ActioPath(parts ...string) string {
	all := append([]string{DirName}, parts...)
	return filepath.Join(all...)
}

// StandardFiles provides common Actio paths by logical key.
// This makes it easy to reference core Actio files without repeating strings.
var StandardFiles = map[string]string{
	"index":       ActioPath(IndexFile),
	"architecture": ActioPath("architecture", "system.md"),
	"interfaces":   ActioPath("interfaces", "contracts.yaml"),
	"patterns":     ActioPath("patterns", "pattern.md"),
	"rules":        ActioPath("rules", "rules.md"),
	"tasks":        ActioPath("tasks", "task.md"),
}
