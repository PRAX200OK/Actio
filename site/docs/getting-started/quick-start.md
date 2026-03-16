# Quick Start

Get Actio running in under a minute.

## 1. Create a new project

From any directory:

```bash
actio create my-app
```

This creates:

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
    └── tasks/
        └── task.md
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

This adds the `actio/` sidecar and `ACTIO.md` without overwriting existing files.

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
