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
	createCmd.Flags().String("preset", "", "Project structure: minimal, standard, or full (default: prompt if TTY, else standard)")
}

var createCmd = &cobra.Command{
	Use:   "create <project_name>",
	Short: "Create a new project with Actio sidecar",
	Long:  "Creates a new directory and Actio sidecar. Use --preset to choose structure (minimal, standard, full); without --preset, prompts interactively when in a TTY.",
	Example: `  actio create my-app
  actio create my-app --preset=minimal
  actio create my-app --preset=full`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("get current directory: %w", err)
		}
		presetStr, _ := cmd.Flags().GetString("preset")
		preset, err := ResolvePreset(presetStr, cmd.ErrOrStderr())
		if err != nil {
			return err
		}
		if err := project.CreateNewProject(cwd, name, preset); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Created Actio-enabled project at %s (preset: %s)\n", filepath.Join(cwd, name), preset.String())
		return nil
	},
}

