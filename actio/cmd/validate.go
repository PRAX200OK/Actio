package cmd

import (
	"fmt"
	"os"

	"actio/internal/validate"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(validateCmd)
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate Actio sidecar structure and configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("get current directory: %w", err)
		}
		issues, err := validate.Validate(cwd)
		if err != nil {
			return fmt.Errorf("validation failed: %w", err)
		}
		if len(issues) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "Actio validation passed")
			return nil
		}
		fmt.Fprintln(cmd.OutOrStdout(), "Actio validation issues:")
		for _, issue := range issues {
			fmt.Fprintf(cmd.OutOrStdout(), "- %s\n", issue)
		}
		return fmt.Errorf("validation failed with %d issue(s)", len(issues))
	},
}

