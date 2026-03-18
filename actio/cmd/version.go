package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// version is the Actio CLI version. For releases, update this or inject via build flags.
const version = "0.1.0"

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolP("short", "s", false, "Print only the version number (e.g. for scripting)")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Actio",
	Long:  "Print the version number of Actio. Use --short to print only the version string.",
	Example: `  actio version
  actio version --short
  actio version -s`,
	RunE: func(cmd *cobra.Command, args []string) error {
		short, _ := cmd.Flags().GetBool("short")
		if short {
			fmt.Fprintln(cmd.OutOrStdout(), version)
			return nil
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Actio v%s\n", version)
		return nil
	},
}