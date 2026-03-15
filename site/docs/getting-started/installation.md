# Installation

Install the Actio CLI to create projects, validate sidecars, and run the MCP server.

## Prerequisites

- **Go 1.22+** (to build from source)
- Or use a pre-built binary when available

## Build from source

Clone the repository and build the `actio` binary:

```bash
git clone https://github.com/PRAX200OK/actio.git
cd actio
go build -o actio ./actio
```

Optionally, move the binary into your `PATH`:

```bash
sudo mv actio /usr/local/bin/
# or
export PATH="$PATH:$(pwd)"
```

Verify:

```bash
actio --help
```

You should see:

```
Actio is an AI-sidecar framework that provides structured context
to AI coding agents to reduce hallucinations and enforce architecture rules.

Usage:
  actio [command]

Available Commands:
  create      Create a new project with Actio sidecar
  doctor      Check Actio project health and print issues
  init        Initialize Actio sidecar in an existing repository
  mcp         Start an MCP-compatible server exposing Actio context
  validate    Validate Actio sidecar structure and configuration
  ...
```

## Next

- [Quick start](/docs/getting-started/quick-start) — run your first commands.
- [Create a project](/docs/getting-started/create-project) — scaffold a full Actio-enabled repo.
