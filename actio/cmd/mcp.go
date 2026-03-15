package cmd

import (
	"fmt"
	"os"

	"actio/internal/mcp"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(mcpCmd)
}

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Start an MCP-compatible server exposing Actio context",
	Long: `Start a minimal Model Context Protocol (MCP) server over stdio.

The server exposes Actio-aware tools for AI coding agents, such as:
- Listing Actio resources (architecture, interfaces, rules, tasks)
- Reading Actio documents
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("get current directory: %w", err)
		}
		return mcp.ServeStdIO(cmd.Context(), cwd, cmd.ErrOrStderr())
	},
}

