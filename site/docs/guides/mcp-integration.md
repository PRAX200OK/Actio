# MCP integration

Actio can expose its context to AI tools via a minimal **MCP** (Model Context Protocol) server. Run `actio mcp` and connect your tool over stdio.

## What is MCP?

MCP is a protocol that lets AI applications talk to external "tools" and "resources." Actio's MCP server provides:

- **Resources** — Actio files (router, architecture, interfaces, rules, tasks, scripts) as readable URIs.
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
      {"uri": "actio://actio/router.yaml", "description": "Actio resource: actio/router.yaml"},
      {"uri": "actio://actio/architecture/system.md", "description": "Actio resource: actio/architecture/system.md"}
    ]
  }
}
```

### mcp.readResource

Returns the content of a resource by URI.

**Request:**

```json
{"jsonrpc":"2.0","id":2,"method":"mcp.readResource","params":{"uri":"actio://actio/router.yaml"}}
```

**Response:**

```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "result": {
    "uri": "actio://actio/router.yaml",
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

## Actio MCP as adapter (plug in other MCPs)

Actio MCP can act as a **single interface** for AI tools (Cursor, Claude Code, etc.): you connect the tool only to Actio, and Actio aggregates its own context and **third-party MCP servers** you configure.

- **Plugin configs** live under **`<repo-root>/mcp/plugins/`**. Add one file per external MCP (e.g. `mcp/plugins/my-tool.yaml`).
- Each file describes how to start that MCP over stdio (command + args, optional env).
- Actio starts those processes, calls their `mcp.listResources` / `mcp.readResource`, and exposes their resources under **`plugin://<name>?uri=...`** alongside built-in **`actio://`** resources.
- The AI tool talks only to `actio mcp`; Actio proxies list/read to the right plugin when the URI is a plugin URI.

Example plugin config (`mcp/plugins/filesystem.yaml`):

```yaml
name: filesystem
description: Local filesystem MCP
command: npx
args: ["-y", "@modelcontextprotocol/server-filesystem", "/path/to/allowed"]
```

After adding configs, restart `actio mcp`. List resources will include both `actio://` and `plugin://` entries; read resource for a plugin URI is proxied to that plugin.

## See also

- [actio mcp](/docs/cli/mcp) — CLI reference.
