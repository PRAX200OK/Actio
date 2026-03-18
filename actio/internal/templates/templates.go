package templates

import (
	"fmt"

	"actio/internal/actio"
)

func ActMD() string {
	return fmt.Sprintf("# Actio Framework\n\n"+
		"This repository uses the **Actio** (AI Context Toolkit) framework.\n\n"+
		"- AI agents **must read** `%s` first before generating or modifying code.\n"+
		"- Architecture and rules defined under `%s/` are the **source of truth**.\n"+
		"- Agents should not modify application code that violates these rules.\n",
		actio.StandardFiles["index"], actio.DirName)
}

const IndexYAML = `version: 1

project:
  name: example-project

domains:
  connectors:
    architecture: architecture/system.md
    interfaces: interfaces/contracts.yaml
    patterns:
      - patterns/pattern.md

rules:
  coding: rules/rules.md

tasks:
  task:
    domain: connectors
    guide: tasks/task.md
`

const ArchitectureSystemMD = `# System Architecture

Describe the high-level system architecture here:

- Core domains
- Bounded contexts
- Allowed dependencies
- Technology choices and constraints
`

const InterfacesContractsYAML = `version: 1

contracts:
  example-connector:
    description: Example connector interface contract
    methods:
      - name: Connect
        input: ConnectionConfig
        output: ConnectionHandle
      - name: Query
        input: QueryRequest
        output: QueryResult
`

const PatternsExampleMD = `# Example Pattern - Snowflake Connector

Intent:
- Standardize how Snowflake connectors are implemented.

Guidelines:
- Use a single connection factory per bounded context.
- Enforce typed query builders over raw SQL where possible.
`

func RulesCodingMD() string {
	return fmt.Sprintf("# Coding Rules\n\n" +
		"- Follow the architecture specified in `%s`.\n" +
		"- Do not introduce new domains without updating `%s`.\n" +
		"- Prefer explicit interfaces documented under `%s`.\n" +
		"- Avoid cross-domain imports that bypass documented boundaries.\n",
		actio.StandardFiles["architecture"],
		actio.StandardFiles["index"],
		actio.ActioPath("interfaces"))
}

func TasksExampleMD() string {
	return fmt.Sprintf("# Example Task - Add New Connector\n\n"+
		"1. Read `%s` to understand the existing domains.\n"+
		"2. Define a new contract in `%s` if needed.\n"+
		"3. Document any patterns in `%s`.\n"+
		"4. Ensure changes comply with rules in `%s`.\n",
		actio.StandardFiles["architecture"],
		actio.StandardFiles["interfaces"],
		actio.ActioPath("patterns"),
		actio.StandardFiles["rules"])
}

// ScriptsManifestYAML returns the content for actio/scripts/manifest.yaml.
// Single maintained file in scripts/: declarative list of runnable scripts plus folder usage.
// Paths are relative to the repository root. Add script entries here when you add new scripts.
const ScriptsManifestYAML = `# actio/scripts/ — single manifest for AI-runnable scripts
# This folder holds scripts the agent can execute. Maintain only this file (manifest.yaml).
# - Python (.py) is the default; add shell (.sh), PowerShell (.ps1), or batch (.bat) as needed; list them below.
# - Agent runs these via the terminal when a task requires it.
# - Keep scripts small, documented, and idempotent where possible.
# - Do not put secrets in scripts; use environment variables or a secrets manager.

version: 1
description: |
  Scripts in actio/scripts/ that can be run from the repo root.
  Use this manifest to discover script paths, descriptions, and usage.

scripts:
  - name: example
    path: actio/scripts/example.py
    description: Example script; confirms the scripts folder is set up.
    usage: Run from repo root: python actio/scripts/example.py (or python3). No arguments. Safe to run anytime.
    when_to_use: When verifying the Actio scripts setup or as a template for new scripts.
`

// ExampleScriptPy returns a minimal example script in Python (default for scripts folder).
const ExampleScriptPy = `#!/usr/bin/env python3
"""Example script – safe for the AI agent to run.
Replace with real automation (e.g. build, lint, codegen)."""
import sys

def main() -> None:
    print("Actio scripts folder is ready. Add your automation here.")
    sys.exit(0)

if __name__ == "__main__":
    main()
`

// ActioPluginsReadmeMD returns the content for actio/plugins/README.md (full preset only).
func ActioPluginsReadmeMD() string {
	return "# Validation plugins\n\n" +
		"Add YAML files here to extend `actio validate` with extra checks.\n\n" +
		"Each file defines a plugin with optional **requiredFiles**. Validation will fail if those files are missing.\n" +
		"See the [plugins guide](https://PRAX200OK.github.io/Actio/docs/guides/plugins) for schema and examples.\n"
}

// MCPPluginsReadmeMD returns the content for mcp/plugins/README.md (full preset only).
func MCPPluginsReadmeMD() string {
	return "# MCP plugin configs\n\n" +
		"Add YAML or JSON files here to plug third-party MCP servers into `actio mcp`.\n\n" +
		"Each file describes how to start one MCP process (command, args, optional env). " +
		"Actio will aggregate their resources under `plugin://` alongside built-in `actio://` resources.\n" +
		"See the [MCP integration guide](https://PRAX200OK.github.io/Actio/docs/guides/mcp-integration) for details.\n"
}

