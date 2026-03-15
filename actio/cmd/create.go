package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"actio/internal/project"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create <project_name>",
	Short: "Create a new project with Actio sidecar",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("get current directory: %w", err)
		}
		if err := project.CreateNewProject(cwd, name); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Created Actio-enabled project at %s\n", filepath.Join(cwd, name))
		return nil
	},
}

