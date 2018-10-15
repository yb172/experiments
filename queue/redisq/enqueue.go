package redisq

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

// DoEnqueue places n messages in redis
func DoEnqueue(queue string, messages int) {
	took, err := enqueue(queue, messages)
	if err != nil {
		log.Fatalf("Error while enqueuing: %v", err)
	}
	log.Printf("It took %v to enqueue %v messages", took, messages)
}

func enqueue(queue string, messages int) (time.Duration, error) {
	client := NewClient()

	msgs := generateMessages(messages)
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
				if err := client.LPush(queue, m).Err(); err != nil {
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
