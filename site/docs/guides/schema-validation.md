# Schema validation

Actio validates `actio/index.yaml` against a **schema** and checks **referential integrity** so the sidecar stays consistent.

## What is validated

### 1. Presence and YAML

- `actio/index.yaml` must exist.
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

- `act/architecture/`, `act/interfaces/`, `act/rules/`, `act/tasks/` must exist as directories.
- Default expected files (e.g. `architecture/system.md`, `rules/coding_rules.md`, `tasks/example_task.md`) are checked if your index references them.

### 5. Plugins

- YAML plugins in `act/plugins/` are loaded; their **requiredFiles** are checked for existence.

## Error messages

Validation returns **human-readable issues**, for example:

- `missing act/index.yaml`
- `index.yaml: version must be set and > 0`
- `index.yaml: project.name must be set`
- `index.yaml: domains.connectors.architecture references missing file: act/architecture/system.md`
- `index.yaml: tasks.deploy.domain references unknown domain: platform`
- `plugin "security": missing required file: docs/SECURITY.md`

Run `act validate` to see all issues; fix paths or add missing files until validation passes.

## See also

- [index.yaml](/docs/concepts/index-yaml)
- [act validate](/docs/cli/validate)
