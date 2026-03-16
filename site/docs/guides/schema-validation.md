# Schema validation

Actio validates `actio/router.yaml` against a **schema** and checks **referential integrity** so the sidecar stays consistent.

## What is validated

### 1. Presence and YAML

- `actio/router.yaml` must exist.
- It must be valid YAML (parseable).

### 2. Required fields

| Field | Rule |
|-------|------|
| **version** | Must be set and &gt; 0. |
| **project.name** | Must be non-empty. |
| **domains.***name*.**architecture** | Required for each domain. |
| **domains.***name*.**interfaces** | Required for each domain. |
| **rules.coding** | Required. |
| **tasks.***name*.**domain** | Required; must match a key in `domains`. |
| **tasks.***name*.**guide** | Required. |

### 3. Referential integrity

- Every path under **domains** (architecture, interfaces, patterns) must point to an existing file under `actio/`.
- **rules.coding** must point to an existing file.
- Every **tasks.***name*.**guide** must point to an existing file.
- Every **tasks.***name*.**domain** must be an existing key in **domains**.

### 4. Directory layout

- `actio/architecture/`, `actio/interfaces/`, `actio/rules/`, `actio/tasks/` must exist as directories.
- Default expected files (e.g. `architecture/system.md`, `rules/rules.md`, `tasks/task.md`) are checked if your router references them.

### 5. Plugins

- YAML plugins in `act/plugins/` are loaded; their **requiredFiles** are checked for existence.

## Error messages

Validation returns **human-readable issues**, for example:

- `missing act/index.yaml`
- `router.yaml: version must be set and > 0`
- `router.yaml: project.name must be set`
- `router.yaml: domains.connectors.architecture references missing file: actio/architecture/system.md`
- `router.yaml: tasks.deploy.domain references unknown domain: platform`
- `plugin "security": missing required file: docs/SECURITY.md`

Run `actio validate` to see all issues; fix paths or add missing files until validation passes.

## See also

- [router.yaml](/docs/concepts/index-yaml)
- [actio validate](/docs/cli/validate)
