package cmd

import (
	"fmt"
	"os"

	"actio/internal/project"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().String("preset", "", "Project structure: minimal, standard, or full (default: prompt if TTY, else standard)")
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Actio sidecar in an existing repository",
	Long:  "Adds Actio sidecar to the current directory. Use --preset to choose structure (minimal, standard, full); without --preset, prompts interactively when in a TTY.",
	Example: `  actio init
  actio init --preset=minimal
  actio init --preset=full`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("get current directory: %w", err)
		}
		presetStr, _ := cmd.Flags().GetString("preset")
		preset, err := ResolvePreset(presetStr, cmd.ErrOrStderr())
		if err != nil {
			return err
		}
		if err := project.InitExistingRepo(cwd, preset); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Initialized Actio sidecar in current repository (preset: %s)\n", preset.String())
		return nil
	},
}

