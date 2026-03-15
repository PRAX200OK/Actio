# Actio

[![Go Report Card](https://goreportcard.com/badge/github.com/PRAX200OK/actio)](https://goreportcard.com/report/github.com/PRAX200OK/actio)
[![GoDoc](https://godoc.org/github.com/PRAX200OK/actio?status.svg)](https://godoc.org/github.com/PRAX200OK/actio)

Actio is an AI-sidecar framework that provides structured context to AI coding agents to reduce hallucinations and enforce architecture rules.

## Installation

### From Source

```bash
go install github.com/PRAX200OK/actio@latest
```

### From Releases

Download the latest binary from the [releases page](https://github.com/PRAX200OK/actio/releases).

## Quick Start

1. Initialize Actio in your project:
```bash
actio init
```

2. Create a new project with Actio:
```bash
actio create my-project
cd my-project
```

3. Validate your Actio setup:
```bash
actio validate
```

## Commands

- `actio init` - Initialize Actio sidecar in an existing repository
- `actio create <name>` - Create a new project with Actio sidecar
- `actio validate` - Validate Actio sidecar structure and configuration
- `actio doctor` - Check project health and print issues
- `actio mcp` - Start an MCP-compatible server exposing Actio context
- `actio version` - Print version information

## Architecture

Actio provides a structured approach to AI context management:

- **actio/index.yaml** - Project index and metadata
- **actio/architecture/** - System architecture documentation
- **actio/interfaces/** - Contract definitions
- **actio/patterns/** - Reusable patterns
- **actio/rules/** - Coding rules and constraints
- **actio/tasks/** - Task-specific guides

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.