package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/codegold79/go-playground/cobra-basic/stv/internal/episode"
)

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Retrieve info from a random episode",
	Long:  `Retrieve info from a random episode from Star Trek Voyager.`,
	Run: func(cmd *cobra.Command, args []string) {
		var ep episode.Details
		ep, err := episode.Random()
		if err != nil {
			fmt.Println("error", err)
		}

		fmt.Println(ep)
	},
}
