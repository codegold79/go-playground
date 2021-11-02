package cmd

import (
	"fmt"
	"strconv"

	"github.com/codegold79/go-playground/cobra-basic/stv/internal/episode"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stv",
	Short: "Provides episode details from number or randomly",
	Long: `Star Trek Voyager Episode Retriever returns brief episode details based
on an episode number. It can also provide information for a random episode.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Print(`Get episode details given an episode number between 1 and 172 inclusive.

Usage:
   stv <episode number>

Use --help for more information on how to use this command.
`)
			return nil
		}

		n, err := validEpisodeNumber(args[0])
		if err != nil {
			return fmt.Errorf("episode number: %w", err)
		}

		ep, err := episode.Info(n)
		if err != nil {
			return err
		}

		fmt.Println(ep)
		return nil
	},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

func validEpisodeNumber(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("got %q but was expecting a number between 1 and 172 inclusive", s)
	}

	if n < 1 || n > 172 {
		return 0, fmt.Errorf("got %d but was expecting a number between 1 and 172 inclusive", n)
	}

	return n, nil
}
