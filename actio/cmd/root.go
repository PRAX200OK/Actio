package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "actio",
	Short: "Actio - AI-sidecar framework for deterministic context",
	Long: `Actio is an AI-sidecar framework that provides structured context
to AI coding agents to reduce hallucinations and enforce architecture rules.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute runs the root command with consistent error handling.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}


