package cmd

import (
	"fmt"
	"path/filepath"

	"actio/internal/validate"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doctorCmd)
	doctorCmd.Flags().StringP("path", "C", ".", "Path to project root (default: current directory)")
}

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check Actio project health and print issues",
	Long:  "Runs the same checks as validate but always exits 0. Use for health checks or CI reporting without failing the build.",
	Example: `  actio doctor
  actio doctor -C /path/to/project`,
	RunE: func(cmd *cobra.Command, args []string) error {
		path, _ := cmd.Flags().GetString("path")
		root, err := filepath.Abs(path)
		if err != nil {
			return fmt.Errorf("resolve path: %w", err)
		}
		issues, err := validate.Validate(root)
		if err != nil {
			return fmt.Errorf("doctor failed: %w", err)
		}
		if len(issues) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "Actio doctor: all checks passed")
			return nil
		}
		fmt.Fprintln(cmd.OutOrStdout(), "Actio doctor found issues:")
		for _, issue := range issues {
			fmt.Fprintf(cmd.OutOrStdout(), "- %s\n", issue)
		}
		return nil
	},
}

