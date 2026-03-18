# Domains

**Domains** in Actio are named areas of your system (e.g. `connectors`, `api`, `billing`). Each domain has its own architecture, interfaces, and patterns so agents get scoped context.

## Definition in router.yaml

```yaml
domains:
  connectors:
    architecture: architecture/system.md
    interfaces: interfaces/contracts.yaml
    patterns:
      - patterns/example_pattern.md
  api:
    architecture: architecture/api.md
    interfaces: interfaces/api.yaml
    patterns: []
```

## Fields per domain

| Field | Required | Description |
|-------|----------|-------------|
| **architecture** | Yes | Path (relative to `actio/`) to the architecture doc for this domain. |
| **interfaces** | Yes | Path to the interfaces/contracts file (e.g. YAML). |
| **patterns** | No | List of pattern doc paths. Can be empty `[]`. |

## Why domains

- **Scoped context** — For “change the connectors domain,” the agent loads only that domain’s architecture, interfaces, and patterns.
- **Bounded contexts** — Maps well to DDD-style boundaries or service boundaries.
- **Task routing** — Tasks reference a domain so the right docs are loaded for “add new connector” vs “add new API endpoint.”

## Example: monorepo

```yaml
domains:
  services-auth:
    architecture: architecture/auth.md
    interfaces: interfaces/auth.yaml
    patterns:
      - patterns/jwt-handling.md
  services-payments:
    architecture: architecture/payments.md
    interfaces: interfaces/payments.yaml
    patterns: []
```

## Validation

- Every path must point to an existing file under `act/`.
- Tasks that reference a domain must use a domain key that exists here.

## Next

- [Rules and tasks](/docs/concepts/rules-and-tasks)
- [Use cases](/docs/guides/use-cases)
