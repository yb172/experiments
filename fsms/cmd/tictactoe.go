package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yb172/experiments/fsms/tictactoe"
)

// gameCmd represents "game" command
var gameCmd = &cobra.Command{
	Use:   "game",
	Short: "Tic Tac Toe game",
}

// startServerCmd starts grpc game server
var startServerCmd = &cobra.Command{
	Use:   "start-server",
	Short: "Start Tic Tac Toe game server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := tictactoe.StartServer(); err != nil {
			log.Fatal(err)
		}
	},
}

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play Tic Tac Toe",
	Run: func(cmd *cobra.Command, args []string) {
		if err := tictactoe.ConnectAndPlay(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	gameCmd.AddCommand(startServerCmd)
	gameCmd.AddCommand(playCmd)
	rootCmd.AddCommand(gameCmd)
}
