package redisq

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
)

// DequeueSimple consumes messages from queue using RPOP
func DequeueSimple(queue string) (time.Duration, error) {
	client := NewClient()

	start := time.Now()

	for {
		err := client.RPop(queue).Err()
		if err != nil {
			if strings.Contains(err.Error(), "redis: nil") {
				break
			}
			return time.Duration(0), fmt.Errorf("error while popping from list: %v", err)
		}
	}

	end := time.Now()
	return end.Sub(start), nil
}

// DequeueSimpleInParallel consumes messages from queue using RPOP
// and "workers" number of concurrent workers
func DequeueSimpleInParallel(queue string) (time.Duration, error) {

	workers := viper.GetInt("workers")
	log.Printf("Start dequeue with RPOP in parallel using %v workers", workers)

	results := make(chan time.Duration)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			took, err := DequeueSimple(queue)
			if err != nil {
				log.Printf("error while running dequeue: %v", err)
			}
			results <- took
			wg.Done()
		}()
	}

	var finalDur time.Duration
	go func() {
		for dur := range results {
			if finalDur.Nanoseconds() < dur.Nanoseconds() {
				finalDur = dur
			}
		}
	}()
	wg.Wait()
	close(results)

	return finalDur, nil
}
