# actio mcp

Start a minimal **MCP-compatible** (Model Context Protocol) server over **stdio**. Exposes Actio resources so AI tools can list and read sidecar files.

## Synopsis

```bash
actio mcp
```

- Runs in the **current directory** (project root).
- Reads JSON-RPC requests from **stdin**, writes responses to **stdout**.
- Logs errors to **stderr**.

## Supported methods

### mcp.listResources

Returns a list of Actio resources (e.g. `router.yaml`, architecture, interfaces, rules, task guides, scripts) that exist under `actio/`.

**Request:**

```json
{"jsonrpc":"2.0","id":1,"method":"mcp.listResources","params":{}}
```

**Response:** `result.resources` is an array of `{ "uri": "actio://actio/...", "description": "..." }`.

### mcp.readResource

Returns the content of a resource by its `actio://` URI.

**Request:**

```json
{"jsonrpc":"2.0","id":2,"method":"mcp.readResource","params":{"uri":"actio://actio/router.yaml"}}
```

**Response:** `result.content` is the file contents (string).

## URIs

Resources use the scheme **actio://** and paths relative to the project root, e.g.:

- `actio://actio/router.yaml`
- `actio://actio/architecture/system.md`
- `actio://actio/rules/rules.md`
- `actio://actio/scripts/manifest.yaml` (scripts manifest: single maintained file — declarative list and usage)

## Example (shell)

```bash
cd my-actio-project
echo '{"jsonrpc":"2.0","id":1,"method":"mcp.listResources","params":{}}' | actio mcp 2>/dev/null
```

## Integration

Configure your AI tool (e.g. Cursor) to run `actio mcp` as an MCP server with transport over stdio. The tool can then call `mcp.listResources` and `mcp.readResource` to load Actio context before generating code.

## See also

- [MCP integration guide](/docs/guides/mcp-integration)
