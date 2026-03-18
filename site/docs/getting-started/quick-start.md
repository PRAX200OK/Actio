# Quick Start

Get Actio running in under a minute.

## 1. Create a new project

From any directory:

```bash
actio create my-app
```

When run in a terminal, you’ll be asked to choose a project structure: **minimal** (core only), **standard** (adds scripts), or **full** (adds scripts and plugin dirs). You can skip the prompt with `--preset=minimal`, `--preset=standard`, or `--preset=full`. See [actio create](/docs/cli/create) for details.

This creates (standard preset):

```
my-app/
├── ENTRYPOINT.yaml
├── src/
└── actio/
    ├── router.yaml
    ├── architecture/
    │   └── system.md
    ├── interfaces/
    │   └── contracts.yaml
    ├── patterns/
    │   └── pattern.md
    ├── rules/
    │   └── rules.md
    ├── tasks/
    │   └── task.md
    └── scripts/
        ├── manifest.yaml   # single maintained file: list scripts and usage
        └── example.py
```

## 2. Validate

```bash
cd my-app
actio validate
```

Expected: `Actio validation passed`.

## 3. Add Actio to an existing repo

If you already have a project:

```bash
cd /path/to/your/repo
actio init
```

You can choose the same structure presets (minimal, standard, full) as with `actio create`. This adds the `actio/` sidecar and `ENTRYPOINT.yaml` without overwriting existing files.

## 4. Check health

```bash
actio doctor
```

Reports any missing or invalid Actio files. Exit code is always 0 (non-blocking).

## What agents see

When an AI agent opens your repo:

1. **ENTRYPOINT.yaml** tells it to read `actio/router.yaml` first.
2. **router.yaml** routes it to architecture, rules, and task guides.
3. The agent uses that context before writing or changing code.

No more “scan everything” — agents get **deterministic context**.

## Next

- [Core concepts: the sidecar](/docs/concepts/sidecar)
- [CLI reference](/docs/cli/create)
