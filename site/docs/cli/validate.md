# actio validate

Check that the Actio sidecar is present, valid, and consistent. Fails with a non-zero exit code if there are issues.

## Synopsis

```bash
actio validate
```

Runs in the **current directory**.

## What it checks

1. **actio/index.yaml** exists and is valid YAML.
2. **Schema:** `version`, `project.name`, `domains`, `rules.coding`, and task references are set and valid.
3. **Referential integrity:** Every file path in `index.yaml` (architecture, interfaces, patterns, rules, task guides) exists under `actio/`.
4. **Required layout:** Directories `actio/architecture/`, `actio/interfaces/`, `actio/rules/`, `actio/tasks/` exist.
5. **Default files:** `architecture/system.md`, `rules/rules.md`, `tasks/task.md` exist (or whatever your index references).
6. **Plugins:** Any YAML plugins under `actio/plugins/` are run; their required files are checked.

## Output

**Success:**

```
Actio validation passed
```

**Failure:** Lists each issue, then exits with non-zero status.

```
Actio validation issues:
- missing actio/index.yaml
- index.yaml: project.name must be set
- index.yaml: tasks.deploy.guide references missing file: actio/tasks/deploy.md
```

## Use in CI

```yaml
# GitHub Actions example
- name: Validate Actio
  run: actio validate
```

Use this to keep the sidecar correct and prevent broken references from being merged.

## See also

- [actio doctor](/docs/cli/doctor) — same checks, always exit 0 (report only).
- [Schema validation](/docs/guides/schema-validation)
