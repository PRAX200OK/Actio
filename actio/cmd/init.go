package cmd

import (
	"fmt"
	"os"

	"actio/internal/project"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Actio sidecar in an existing repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("get current directory: %w", err)
		}
		if err := project.InitExistingRepo(cwd); err != nil {
			return err
		}
		fmt.Fprintln(cmd.OutOrStdout(), "Initialized Actio sidecar in current repository")
		return nil
	},
}

