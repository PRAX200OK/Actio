# actio doctor

Run the same checks as [actio validate](/docs/cli/validate) but **always exit 0**. Use it for a quick health check without failing scripts or CI.

## Synopsis

```bash
actio doctor
```

Runs in the **current directory**.

## Output

**All OK:**

```
Actio doctor: all checks passed
```

**Issues found:** Prints the same issue list as `actio validate`, but the process still exits with code 0.

```
Actio doctor found issues:
- missing actio/index.yaml
- index.yaml: rules.coding references missing file: actio/rules/rules.md
```

## When to use

- **Local health check** — “Is my ACT sidecar in good shape?”
- **Non-blocking CI** — Report ACT issues in logs without failing the build.
- **Onboarding** — New contributors can run `act doctor` to see what’s missing.

## See also

- [actio validate](/docs/cli/validate) — strict validation with non-zero exit on failure.
