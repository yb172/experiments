package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yb172/experiments/queue/redisq"
)

// redisCmd represents the redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "Run queue experiments with redis",
}

var enqueueCmd = &cobra.Command{
	Use:   "enqueue",
	Short: "Enqueue 1M messages",
	Run: func(cmd *cobra.Command, args []string) {
		took, err := redisq.Enqueue()
		if err != nil {
			log.Fatalf("Error while enqueuing: %v", err)
		}
		log.Printf("It took %v to enqueue 1M messages", took)
	},
}

var dequeueCmd = &cobra.Command{
	Use:   "dequeue",
	Short: "Dequeue 1M messages",
}

var dequeueSimpleCmd = &cobra.Command{
	Use:   "pop",
	Short: "Using RPOP",
	Run: func(cmd *cobra.Command, args []string) {
		took, err := redisq.DequeueSimple()
		if err != nil {
			log.Fatalf("Error while dequeuing: %v", err)
		}
		log.Printf("It took %v to dequeue 1M messages", took)
	},
}

var dequeueSimpleInParallelCmd = &cobra.Command{
	Use:   "pop-parallel",
	Short: "Using RPOP and multiple workers",
	Run: func(cmd *cobra.Command, args []string) {
		took, err := redisq.DequeueSimpleInParallel()
		if err != nil {
			log.Fatalf("Error while dequeuing: %v", err)
		}
		log.Printf("It took %v to dequeue 1M messages", took)
	},
}

func init() {
	dequeueCmd.AddCommand(dequeueSimpleCmd)
	dequeueCmd.AddCommand(dequeueSimpleInParallelCmd)

	redisCmd.AddCommand(enqueueCmd)
	redisCmd.AddCommand(dequeueCmd)

	rootCmd.AddCommand(redisCmd)
}
