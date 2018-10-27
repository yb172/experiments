package redisq

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
)

// DequeueReliably consumes messages from queue using RPOPLPUSH
func DequeueReliably(queue string) (time.Duration, error) {
	client := NewClient()

	start := time.Now()
	pqueue := fmt.Sprintf("%s:processing", queue)

	received := 0
	delay := viper.GetDuration("delay")
	for {
		res, err := client.RPopLPush(queue, pqueue).Result()
		if err != nil {
			if strings.Contains(err.Error(), "redis: nil") {
				break
			}
			return time.Duration(0), fmt.Errorf("error while popping from list: %v", err)
		}

		// Simulate delay
		time.Sleep(delay)
		received++

		err = client.LRem(pqueue, 1, res).Err()
		if err != nil {
			return time.Duration(0), fmt.Errorf("error while removing item %v from processing queue: %v", res, err)
		}
	}

	end := time.Now()
	// Remove artificial delay
	for i := 0; i < received; i++ {
		end = end.Add(-1 * delay)
	}
	return end.Sub(start), nil
}

// DequeueReliablyInParallel consumes messages from queue using RPOPLPUSH
// and "workers" number of concurrent workers
func DequeueReliablyInParallel(queue string) (time.Duration, error) {
	workers := viper.GetInt("workers")
	log.Printf("Start dequeue with RPOPLPUSH & LREM in parallel using %v workers", workers)

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
