package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yb172/experiments/kube/testtools/load"
)

func init() {
	rootCmd.AddCommand(loadCmd)
}

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Generate some load to the service",
	Long:  `Program periodically tirggers application endpoint`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := load.GenerateLoad(); err != nil {
			log.Fatalf("Error while generating load: %v", err)
		}
	},
}
