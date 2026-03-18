# actio init

Add the Actio sidecar to an **existing** repository. Does not overwrite existing files. You choose how much structure to generate via **presets**.

## Synopsis

```bash
actio init [--preset=minimal|standard|full]
```

Runs in the **current directory**. Creates `actio/` and its default files, and adds `ENTRYPOINT.yaml` only if it does not already exist.

- **--preset** — Optional. `minimal`, `standard`, or `full`. If omitted and stdin is a TTY, you are prompted; otherwise `standard` is used.

## Presets

| Preset | What you get |
|--------|----------------|
| **minimal** | Core only: no `actio/scripts/`. |
| **standard** | Core plus `actio/scripts/` (manifest.yaml + example.py). Default. |
| **full** | Standard plus `actio/plugins/` and `mcp/plugins/` with READMEs. |

## When to use

- You already have a repo and want to adopt Actio.
- You want a sidecar layout; choose minimal, standard, or full depending on whether you need scripts and plugin dirs.

## Behavior

1. Creates `actio/` and subdirs according to the preset (see [actio create](/docs/cli/create) for preset details).
2. Writes default files only where they are **missing**.
3. Writes `ENTRYPOINT.yaml` only if it does not exist (does not overwrite your custom `ENTRYPOINT.yaml`).

## Examples

```bash
cd /path/to/existing/repo
actio init
# If TTY: prompts for preset. Then: Initialized Actio sidecar in current repository (preset: standard)

actio init --preset=minimal
# No prompt; minimal layout only.
```

## See also

- [actio create](/docs/cli/create) — for a new project from scratch (same presets).
- [Quick start](/docs/getting-started/quick-start)
