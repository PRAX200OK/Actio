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
	return fmt.Sprintf("# Example Task - Add New Connector\n\n" +
		"1. Read `%s` to understand the existing domains.\n" +
		"2. Define a new contract in `%s` if needed.\n" +
		"3. Document any patterns in `%s`.\n" +
		"4. Ensure changes comply with rules in `%s`.\n",
		actio.StandardFiles["architecture"],
		actio.StandardFiles["interfaces"],
		actio.ActioPath("patterns"),
		actio.StandardFiles["rules"])
}

