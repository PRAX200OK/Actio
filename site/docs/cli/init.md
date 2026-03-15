# actio init

Add the Actio sidecar to an **existing** repository. Does not overwrite existing files.

## Synopsis

```bash
actio init
```

Runs in the **current directory**. Creates `actio/` and its default files, and adds `ACTIO.md` only if it does not already exist.

## When to use

- You already have a repo and want to adopt Actio.
- You want the standard sidecar layout and example content without creating a new folder.

## Behavior

1. Creates `actio/` and subdirs: `architecture/`, `interfaces/`, `patterns/`, `rules/`, `tasks/`.
2. Writes default files only where they are **missing** (e.g. `index.yaml`, `architecture/system.md`, …).
3. Writes `ACTIO.md` only if it does not exist (does not overwrite your custom `ACTIO.md`).

## Example

```bash
cd /path/to/existing/repo
actio init
# Initialized Actio sidecar in current repository
```

## See also

- [act create](/docs/cli/create) — for a new project from scratch.
- [Quick start](/docs/getting-started/quick-start)
