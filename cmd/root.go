package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "cliChatApp",
		Long:  `Create a CLI chat application using golang.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
