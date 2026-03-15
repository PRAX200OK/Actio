# Create a project

Use `actio create` to scaffold a new project with the full Actio sidecar.

## Command

```bash
actio create <project_name>
```

- **project_name** — Name of the new directory (e.g. `my-service`, `api-gateway`).
- The project is created in the **current working directory**.

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
├── ACTIO.md              # Entry file for AI agents
├── src/                # Your application code (empty)
└── actio/
    ├── index.yaml      # Context router
    ├── architecture/
    │   └── system.md   # System architecture doc
    ├── interfaces/
    │   └── contracts.yaml
    ├── patterns/
    │   └── example_pattern.md
    ├── rules/
    │   └── rules.md
    └── tasks/
        └── task.md
```

## If the directory already exists

```bash
actio create existing-dir
# Error: directory /path/to/existing-dir already exists
```

Choose another name or use [actio init](/docs/cli/init) to add Actio inside an existing repo.

## Next

- [Understanding index.yaml](/docs/concepts/index-yaml)
- [Init in an existing repo](/docs/cli/init)
