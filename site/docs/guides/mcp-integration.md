# MCP integration

Actio can expose its context to AI tools via a minimal **MCP** (Model Context Protocol) server. Run `actio mcp` and connect your tool over stdio.

## What is MCP?

MCP is a protocol that lets AI applications talk to external "tools" and "resources." Actio's MCP server provides:

- **Resources** — Actio files (index, architecture, rules, tasks) as readable URIs.
- **List + read** — The client can list available resources and read their contents.

## Starting the server

From your project root:

```bash
actio mcp
```

The process reads JSON-RPC from stdin and writes responses to stdout. It does not exit unless stdin closes or an error occurs.

## Methods

### mcp.listResources

Returns all known Actio resources (files that exist under `actio/`).

**Request:**

```json
{"jsonrpc":"2.0","id":1,"method":"mcp.listResources","params":{}}
```

**Response (example):**

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "resources": [
      {"uri": "actio://actio/index.yaml", "description": "Actio resource: actio/index.yaml"},
      {"uri": "actio://actio/architecture/system.md", "description": "Actio resource: actio/architecture/system.md"}
    ]
  }
}
```

### mcp.readResource

Returns the content of a resource by URI.

**Request:**

```json
{"jsonrpc":"2.0","id":2,"method":"mcp.readResource","params":{"uri":"actio://actio/index.yaml"}}
```

**Response:**

```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "result": {
    "uri": "actio://actio/index.yaml",
    "content": "version: 1\nproject:\n  name: ..."
  }
}
```

## Cursor and other tools

To use Actio with Cursor (or another MCP client):

1. Configure the client to run a **stdio** MCP server.
2. Set the command to `actio mcp` and the working directory to your project root.
3. The client can then call `mcp.listResources` and `mcp.readResource` to inject Actio context into the AI session.

Exact steps depend on the client’s MCP support; refer to its docs for “MCP server” or “stdio transport.”

## See also

- [actio mcp](/docs/cli/mcp) — CLI reference.
