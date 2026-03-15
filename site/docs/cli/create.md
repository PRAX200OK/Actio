# actio create

Create a new project with the full Actio sidecar and `ACTIO.md` entry file.

## Synopsis

```bash
actio create <project_name>
```

## Arguments

| Argument | Required | Description |
|----------|----------|-------------|
| **project_name** | Yes | Name of the new project directory. |

The project is created in the **current working directory**.

## Example

```bash
actio create my-service
# Created Actio-enabled project at /path/to/cwd/my-service
```

## Generated layout

- **&lt;project_name&gt;/ACTIO.md**
- **&lt;project_name&gt;/src/** (empty)
- **&lt;project_name&gt;/actio/** with:
  - `index.yaml`
  - `architecture/system.md`
  - `interfaces/contracts.yaml`
  - `patterns/example_pattern.md`
  - `rules/coding_rules.md`
  - `tasks/example_task.md`

## Errors

- **Directory already exists** — Choose a different name or use [actio init](/docs/cli/init) to add Actio inside an existing directory.

## See also

- [Getting started: Create a project](/docs/getting-started/create-project)
- [actio init](/docs/cli/init)
