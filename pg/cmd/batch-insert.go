package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yb172/experiments/pg/batchinsert"
)

var count int

// batchInsertCmd represents the init command
var batchInsertCmd = &cobra.Command{
	Use:   "batch-insert",
	Short: "See how fast batch-insert is happening",
	Long:  `How long would it take to batch-insert 100K records?`,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	if err := batchinsert.RunBatchInsert(count); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(batchInsertCmd)
	batchInsertCmd.Flags().IntVarP(&count, "count", "c", 100000, "# of records to insert")
}
