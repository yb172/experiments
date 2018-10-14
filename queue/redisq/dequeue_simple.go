package redisq

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// DequeueSimple consumes messages from queue using RPOP
func DequeueSimple() (time.Duration, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	start := time.Now()
	log.Println("Start dequeue")

	for {
		err := client.RPop("queue").Err()
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
func DequeueSimpleInParallel() (time.Duration, error) {

	workers := viper.GetInt("workers")
	log.Printf("Start dequeue in parallel using %v workers", workers)

	results := make(chan time.Duration)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			took, err := DequeueSimple()
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
