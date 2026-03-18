# Rules and tasks

Actio separates **rules** (what must always be true) from **tasks** (how to do a specific thing). Both are referenced from `actio/router.yaml`.

## Rules

Rules are global constraints agents must follow. The only required one in the schema is **coding**.

### rules.coding

Points to a single file (e.g. `rules/coding_rules.md`) that describes:

- Architecture boundaries and allowed dependencies
- Naming and style expectations
- What not to do (e.g. no cross-domain imports without a documented interface)

Example content:

```markdown
# Coding Rules

- Follow the architecture in `actio/architecture/system.md`.
- Do not introduce new domains without updating `actio/router.yaml`.
- Prefer interfaces documented under `actio/interfaces/`.
- Avoid cross-domain imports that bypass documented boundaries.
```

Agents should load this file whenever they are about to generate or modify code.

## Tasks

Tasks are **named procedures** tied to a domain and a guide.

### Structure

```yaml
tasks:
  add_connector:
    domain: connectors
    guide: tasks/add_connector.md
  deploy_staging:
    domain: platform
    guide: tasks/deploy_staging.md
```

- **domain** — Must match a key under `domains`. Provides which architecture/interfaces/patterns apply.
- **guide** — Path to a markdown file (relative to `actio/`) that describes the steps.

### How agents use tasks

1. User says: “Do the add_connector task.”
2. Agent reads `actio/router.yaml`, finds task `add_connector`.
3. Agent loads the **domain** `connectors` (architecture, interfaces, patterns).
4. Agent reads the **guide** `tasks/add_connector.md`.
5. Agent follows the guide and respects the domain context.

This gives **minimal, deterministic context** per task instead of “read the whole repo.”

### Example guide

**actio/tasks/add_connector.md:**

```markdown
# Add New Connector

1. Read `actio/architecture/system.md` for existing domains.
2. Add a contract in `actio/interfaces/contracts.yaml` if needed.
3. Document the pattern in `actio/patterns/` if reusable.
4. Implement following `actio/rules/coding_rules.md`.
```

## Validation

- **rules.coding** must be set and the file must exist.
- Every task’s **domain** must exist in **domains**.
- Every task’s **guide** file must exist.

Run `actio validate` to check.

## Next

- [CLI: validate and doctor](/docs/cli/validate)
- [Plugins](/docs/guides/plugins)
