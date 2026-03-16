# Introduction

**Actio** (AI Context Toolkit) is an AI-sidecar framework that gives AI coding agents (Cursor, Copilot, etc.) **structured context** so they follow your architecture and rules instead of guessing.

## The problem

AI agents hallucinate or drift because:

- **Context is unstructured** — they don't know which files matter.
- **Architecture isn't explicit** — rules live in heads or scattered docs.
- **There's no single entry point** — agents scan randomly.

## The solution

Actio adds a small **sidecar** to your repo: an `actio/` directory and an `ENTRYPOINT.yaml` entry file. The sidecar is the **source of truth** for agents.

- **Deterministic entry point:** Agents read `actio/router.yaml` first.
- **Structured routing:** The index points to architecture, interfaces, patterns, rules, and task guides.
- **Validation:** The CLI checks that the sidecar is complete and consistent.

## What you get

| Feature | Description |
|--------|-------------|
| **Sidecar layout** | `actio/` with `architecture/`, `interfaces/`, `patterns/`, `rules/`, `tasks/` |
| **Index router** | `router.yaml` routes agents to the right docs per domain and task |
| **CLI** | `actio create`, `actio init`, `actio validate`, `actio doctor`, `actio mcp` |
| **Schema validation** | Index and referenced files are validated |
| **Plugins** | YAML plugins under `actio/plugins/` for extra checks |
| **MCP server** | Expose Actio context to AI tools over stdio |

## Who it's for

- Teams using **Cursor**, **GitHub Copilot**, or other AI coding tools.
- Repos that need **consistent architecture** and **clear boundaries**.
- Anyone who wants **less hallucination** and **more predictable** AI-assisted code.

## Next steps

- [Install the CLI](/docs/getting-started/installation)
- [Create your first Actio project](/docs/getting-started/create-project)
- [Understand the sidecar and index](/docs/concepts/sidecar)
