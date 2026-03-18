# actio create

Create a new project with an Actio sidecar and `ENTRYPOINT.yaml`. You choose how much structure to generate via **presets**.

## Synopsis

```bash
actio create <project_name> [--preset=minimal|standard|full]
```

- **project_name** — Name of the new project directory (created in the current working directory).
- **--preset** — Optional. `minimal`, `standard`, or `full`. If omitted and stdin is a TTY, you are prompted; otherwise `standard` is used (e.g. in CI).

## Presets

| Preset | What you get |
|--------|----------------|
| **minimal** | Core only: `router.yaml`, `architecture/`, `interfaces/`, `patterns/`, `rules/`, `tasks/`, and their default files. No `scripts/`. Best for small repos or CI. |
| **standard** | Minimal plus `actio/scripts/` (single file `manifest.yaml` + `example.py`). Default when not specified. |
| **full** | Standard plus `actio/plugins/` and `mcp/plugins/` with READMEs (for validation plugins and MCP adapter configs). |

## Examples

```bash
actio create my-service
# If in a TTY: prompts to choose project structure (minimal / standard / full)
# Created Actio-enabled project at /path/to/cwd/my-service (preset: standard)

actio create my-service --preset=minimal
# No prompt; creates minimal layout.

actio create my-service --preset=full
# Creates full layout including actio/plugins and mcp/plugins.
```

## Generated layout (standard preset)

- **&lt;project_name&gt;/ENTRYPOINT.yaml**
- **&lt;project_name&gt;/src/** (empty)
- **&lt;project_name&gt;/actio/** with:
  - `router.yaml`
  - `architecture/system.md`
  - `interfaces/contracts.yaml`
  - `patterns/pattern.md`
  - `rules/rules.md`
  - `tasks/task.md`
  - `scripts/manifest.yaml`, `scripts/example.py`

With **minimal**, `actio/scripts/` is omitted. With **full**, `actio/plugins/` and `mcp/plugins/` (and their READMEs) are added.

## Errors

- **Directory already exists** — Choose a different name or use [actio init](/docs/cli/init) to add Actio inside an existing directory.
- **Invalid preset** — `--preset` must be `minimal`, `standard`, or `full`.

## See also

- [Getting started: Create a project](/docs/getting-started/create-project)
- [actio init](/docs/cli/init)
