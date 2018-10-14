package redisq

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

// Message represents avg message
type Message struct {
	ID          int64             `json:"id"`
	Name        string            `json:"name"`
	Ref         string            `json:"ref"`
	Description string            `json:"string"`
	Params      map[string]string `json:"params"`
}

// generateMessages generates 1M random messages
func generateMessages() []Message {
	var msgs []Message
	for i := 0; i < 1000000; i++ {
		msgs = append(msgs, Message{
			ID:          rand.Int63(),
			Name:        "The name of the job. Could be anything",
			Ref:         "Reference to some root / parent object",
			Description: "Message to test queues",
			Params: map[string]string{
				"Attempts":     "18",
				"State":        "Colorado",
				"MaxDeviation": "46",
			},
		})
	}
	return msgs
}

func encodeMessages(msgs []Message) ([][]byte, error) {
	var encoded [][]byte
	for _, m := range msgs {
		e, err := json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("error while encoding message: %v", err)
		}
		encoded = append(encoded, e)
	}
	return encoded, nil
}
