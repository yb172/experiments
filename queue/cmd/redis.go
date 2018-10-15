package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yb172/experiments/queue/redisq"
)

// redisCmd represents the redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "Run queue experiments with redis",
}

var dequeueSimpleCmd = &cobra.Command{
	Use:   "simple",
	Short: "Enqueue and then dequeue using RPOP and multiple workers",
	Long:  "Dequeue by multiple workers",
	Run: func(cmd *cobra.Command, args []string) {
		queueName := time.Now().Format(time.RFC822)
		messagesCount := viper.GetInt("messages")

		redisq.DoEnqueue(queueName, messagesCount)
		took, err := redisq.DequeueSimpleInParallel(queueName)
		if err != nil {
			log.Fatalf("Error while dequeuing: %v", err)
		}
		log.Printf("It took %v to dequeue %v messages", took, messagesCount)
	},
}

func init() {
	redisCmd.AddCommand(dequeueSimpleCmd)

	rootCmd.AddCommand(redisCmd)
}
