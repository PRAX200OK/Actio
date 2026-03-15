# Plugins

Plugins extend Actio validation with **extra checks** (e.g. "this file must exist"). They are YAML files under `actio/plugins/`.

## Plugin layout

```
actio/
└── plugins/
    ├── security.yaml
    └── docs.yaml
```

Only `.yaml` and `.yml` files are loaded. Subdirectories are ignored.

## Plugin schema

Each plugin file can define:

| Field | Required | Description |
|-------|----------|-------------|
| **name** | No | Display name (defaults to filename). |
| **description** | No | Short description. |
| **requiredFiles** | No | List of paths (relative to repo root) that must exist. |

## Example

**act/plugins/security.yaml:**

```yaml
name: security-guardrails
description: Ensure security docs and config exist
requiredFiles:
  - act/rules/coding_rules.md
  - docs/SECURITY.md
  - .github/dependabot.yml
```

When you run `act validate` or `act doctor`, the CLI:

1. Scans `act/plugins/*.yaml`.
2. For each plugin, checks that every path in `requiredFiles` exists.
3. Reports: `plugin "security-guardrails": missing required file: docs/SECURITY.md` for any missing path.

## Use cases

- **Compliance:** Require ADRs, SECURITY.md, or policy docs.
- **Monorepo:** Require per-package README or config.
- **Custom rules:** Require files that your team agreed on (e.g. `docs/architecture/adr-001.md`).

## No plugins directory

If `act/plugins/` does not exist, no plugin checks run. Validation still runs for the core ACT layout and index.

## See also

- [act validate](/docs/cli/validate)
- [Schema validation](/docs/guides/schema-validation)
