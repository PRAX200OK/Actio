# index.yaml

`actio/index.yaml` is the **context router**. It tells agents which files to read for which domain, rule, or task.

## Structure

```yaml
version: 1

project:
  name: my-project

domains:
  connectors:
    architecture: architecture/system.md
    interfaces: interfaces/contracts.yaml
    patterns:
      - patterns/example_pattern.md

rules:
  coding: rules/coding_rules.md

tasks:
  example_task:
    domain: connectors
    guide: tasks/example_task.md
```

## Top-level fields

| Field | Required | Description |
|-------|----------|-------------|
| **version** | Yes | Schema version (must be `1` or higher). |
| **project.name** | Yes | Project name. |
| **domains** | Yes | Named domains; each has architecture, interfaces, and optional patterns. |
| **rules** | Yes | At least `rules.coding` pointing to a rules file. |
| **tasks** | No | Named tasks; each has a `domain` and a `guide` file. |

## Domains

Each **domain** (e.g. `connectors`, `api`, `frontend`) has:

- **architecture** — One file (e.g. `architecture/system.md` or per-domain).
- **interfaces** — One file (e.g. `interfaces/contracts.yaml`).
- **patterns** — List of pattern docs (e.g. `patterns/snowflake.md`).

Paths are relative to `actio/`. The CLI validates that these files exist.

## Rules

- **rules.coding** — Path to the main coding/architecture rules file (e.g. `rules/coding_rules.md`).

## Tasks

Each **task** links a **domain** to a **guide**:

- **domain** — Must match a key under `domains`.
- **guide** — Path to a markdown file (e.g. `tasks/example_task.md`) that describes how to do that task.

Agents can use tasks to load minimal context: “For task X, read this guide and the domain’s architecture/interfaces.”

## Validation

Running `actio validate` checks:

- `index.yaml` exists and is valid YAML.
- **version** and **project.name** are set.
- Every referenced file exists under `actio/`.
- Every task’s **domain** exists in **domains**.
- Required directories (`architecture/`, `interfaces/`, `rules/`, `tasks/`) exist.

## Next

- [Domains in depth](/docs/concepts/domains)
- [Schema validation](/docs/guides/schema-validation)
