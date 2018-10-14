package redisq

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// Enqueue places 1M messages in redis
func Enqueue() (time.Duration, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	msgs := generateMessages()
	emsgs, err := encodeMessages(msgs)
	log.Println("Start enqueue")
	if err != nil {
		return time.Duration(0), fmt.Errorf("error while encoding: %v", err)
	}

	in := make(chan []byte)
	go func() {
		for _, m := range emsgs {
			in <- m
		}
		close(in)
	}()

	var wg sync.WaitGroup
	workers := viper.GetInt("workers")
	start := time.Now()

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			for m := range in {
				if err := client.LPush("queue", m).Err(); err != nil {
					log.Printf("error while pushing to list: %v", err)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	end := time.Now()
	return end.Sub(start), nil
}
