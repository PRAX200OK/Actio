package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Actio",
	Long:  `Print the version number of Actio`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Actio v%s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}