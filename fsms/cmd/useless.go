package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/yb172/experiments/fsms/useless"
)

// uselessCmd represents "fsms useless" command
var uselessCmd = &cobra.Command{
	Use:   "useless",
	Short: "Most useless machine... ever!",
	Run: func(cmd *cobra.Command, args []string) {
		um := useless.Create()

		fmt.Printf("\nEnter commands: on / off\n")
		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			switch text {
			case "on\n":
				fmt.Printf("\nTurning on...\n")
				if err := um.FSM.Event("turn-on"); err != nil {
					log.Printf("Error happened: %v", err)
				}
			case "off\n":
				fmt.Printf("\nTurning off...\n")
				if err := um.FSM.Event("turn-off"); err != nil {
					log.Printf("Error happened: %v", err)
				}
			default:
				fmt.Println("?")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(uselessCmd)
}
