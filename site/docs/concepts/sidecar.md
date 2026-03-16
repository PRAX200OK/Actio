# The Actio sidecar

The **sidecar** is a directory named `actio/` at your project root. It holds all context that AI agents are supposed to read before changing code.

## Why "sidecar"

The app lives in `src/` (or your usual layout); the **actio/** directory sits beside it and describes it. Agents are instructed to read the sidecar first, then apply that context when editing the main codebase.

## Layout

```
your-repo/
├── ENTRYPOINT.yaml    # Entry file: "read actio/router.yaml first"
├── src/             # Your code
└── actio/             # Sidecar
    ├── router.yaml  # Router: points to everything below
    ├── architecture/
    ├── interfaces/
    ├── patterns/
    ├── rules/
    ├── tasks/
    └── plugins/     # Optional
```

## Responsibilities

| Path | Purpose |
|------|--------|
| **actio/router.yaml** | Single entry for agents; routes to domains, rules, tasks |
| **actio/architecture/** | High-level system design, bounded contexts, tech choices |
| **actio/interfaces/** | Contracts (APIs, connectors, boundaries) |
| **actio/patterns/** | Reusable patterns (e.g. "Snowflake connector", "event handler") |
| **actio/rules/** | Coding and architectural rules agents must follow |
| **actio/tasks/** | Task guides (e.g. "add new connector", "deploy") |
| **actio/plugins/** | Optional YAML plugins for extra validation |

## Contract with agents

**ENTRYPOINT.yaml** at the repo root states:

- AI agents **must read** `actio/router.yaml` first.
- Content under `actio/` is the **source of truth**.
- Agents must not violate the architecture and rules defined there.

So the sidecar is not “nice to have” docs — it’s the **contract** for AI-assisted changes.

## Next

- [index.yaml: the router](/docs/concepts/index-yaml)
- [Domains, rules, and tasks](/docs/concepts/rules-and-tasks)
