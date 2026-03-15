package cmd

import (
	"fmt"
	"os"

	"actio/internal/validate"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doctorCmd)
}

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check Actio project health and print issues",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("get current directory: %w", err)
		}
		issues, err := validate.Validate(cwd)
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

