package cmd

import (
	"fmt"
	"path/filepath"

	"actio/internal/validate"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().StringP("path", "C", ".", "Path to project root (default: current directory)")
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate Actio sidecar structure and configuration",
	Long:  "Checks router.yaml, required dirs/files, and referential integrity. Exits non-zero if there are issues.",
	Example: `  actio validate
  actio validate -C /path/to/project
  actio validate --path ./my-app`,
	RunE: func(cmd *cobra.Command, args []string) error {
		path, _ := cmd.Flags().GetString("path")
		root, err := filepath.Abs(path)
		if err != nil {
			return fmt.Errorf("resolve path: %w", err)
		}
		issues, err := validate.Validate(root)
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

