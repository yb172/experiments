package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yb172/experiments/fsms/useless"
)

// uselessCmd represents the redis command
var uselessCmd = &cobra.Command{
	Use:   "useless",
	Short: "Most useless machine... ever!",
	Run: func(cmd *cobra.Command, args []string) {
		um := useless.Create()

		fmt.Printf("\nTurning on...\n")
		if err := um.FSM.Event("turn-on"); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\nTurning off...\n")
		if err := um.FSM.Event("turn-off"); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(uselessCmd)
}
