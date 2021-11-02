package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stv",
	Short: "Provides episode name from number",
	Long: `Star Trek Voyager Episode Name Retriever provides 
the episode title based on an episode number.`,
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
}
