package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yb172/experiments/kube/testtools/load"
)

var attachKeyboard bool

func init() {
	loadCmd.Flags().BoolVarP(&attachKeyboard, "keyboard", "k", true, "Allow dynamic configuration from keyboard")
	rootCmd.AddCommand(loadCmd)
}

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Generate some load to the service",
	Long:  `Program periodically tirggers application endpoint`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := load.GenerateLoad(attachKeyboard); err != nil {
			log.Fatalf("Error while generating load: %v", err)
		}
	},
}
