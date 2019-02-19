package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gen-dev",
	Short: "gen-dev is a tool for local gen development",
	Long:  `gen-dev helps to do things that make local development simpler, e.g. add some load to the system`,
}

// Execute runs the program
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error while running program: %v", err)
	}
}
