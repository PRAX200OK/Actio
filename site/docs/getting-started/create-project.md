# Create a project

Use `actio create` to scaffold a new project with an Actio sidecar. You can choose the structure: **minimal** (core only), **standard** (adds scripts), or **full** (adds scripts and plugin dirs). Without `--preset`, you’ll be prompted in a TTY; in CI, use `--preset=standard` (or another) to avoid prompts.

## Command

```bash
actio create <project_name> [--preset=minimal|standard|full]
```

- **project_name** — Name of the new directory (e.g. `my-service`, `api-gateway`).
- The project is created in the **current working directory**.
- See [actio create](/docs/cli/create) for preset details.

## Example

```bash
actio create demo-api
```

Output:

```
Created Actio-enabled project at /path/to/cwd/demo-api
```

## Generated structure

```
<project_name>/
├── ENTRYPOINT.yaml       # Entry file for AI agents
├── src/                # Your application code (empty)
└── actio/
    ├── router.yaml     # Context router
    ├── architecture/
    │   └── system.md   # System architecture doc
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

## If the directory already exists

```bash
actio create existing-dir
# Error: directory /path/to/existing-dir already exists
```

Choose another name or use [actio init](/docs/cli/init) to add Actio inside an existing repo.

## Next

- [Understanding router.yaml](/docs/concepts/index-yaml)
- [Init in an existing repo](/docs/cli/init)
