# Use cases

Actio is **use-case agnostic**. The same sidecar layout and index work for many kinds of projects. Below are common patterns.

## Microservices / multi-domain apps

- **Domains** = services or bounded contexts (e.g. `auth`, `payments`, `notifications`).
- **Architecture** per domain = service boundaries, allowed dependencies.
- **Interfaces** = API contracts or event schemas between services.
- **Tasks** = “add new service,” “add endpoint to X,” “change contract Y.”

## Data / ML pipelines

- **Domains** = stages (e.g. `ingestion`, `transform`, `model`, `serving`).
- **Interfaces** = schemas or contracts between stages.
- **Patterns** = “add new source,” “schema evolution.”
- **Tasks** = “add new source,” “change schema,” “retrain model.”

## Monorepos

- **Domains** = apps or packages (e.g. `apps/web`, `packages/auth`).
- **Architecture** = dependency rules, package boundaries.
- **Tasks** = “add package,” “bump dependency,” “add app.”

## Compliance and audit

- **Rules** = security and compliance rules (e.g. “no secrets in code,” “all APIs documented”).
- **Plugins** = require `SECURITY.md`, ADRs, or policy docs.
- **Tasks** = “security review,” “audit checklist,” “remediation steps.”

## API-first / design-first

- **Interfaces** = OpenAPI or contract files.
- **Patterns** = API design guidelines (versioning, errors, auth).
- **Tasks** = “add endpoint,” “version API,” “update contract.”

## Incident and ops runbooks

- **Tasks** = runbooks (e.g. “database incident,” “rollback,” “scale up”).
- **Guides** = step-by-step markdown.
- **Rules** = operational constraints (e.g. “no direct prod access from scripts”).
- **Scripts** = put automation in `actio/scripts/` so the agent can run build, lint, or codegen when needed.

---

Pick domains, tasks, and scripts that match your workflow; the framework only enforces structure and references, not the content.
