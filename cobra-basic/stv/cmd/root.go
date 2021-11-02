package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stv",
	Short: "Provides episode details from number or randomly",
	Long: `Star Trek Voyager Episode Retriever returns brief episode details based
on an episode number. It can also provide information for a random episode.`,
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(randomCmd)
}
