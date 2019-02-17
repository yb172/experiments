package gen

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/yb172/experiments/kube/gateway/rpc"
)

// GenerateSeq generates sequence
func GenerateSeq() (string, error) {
	// Seed rand
	rand.Seed(time.Now().UTC().UnixNano())

	i := rand.Intn(4)
	var parts []string
	for i != 0 {
		switch i {
		case 1:
			word, err := rpc.GenerateWord()
			if err != nil {
				return "", fmt.Errorf("error while calling service: %v", err)
			}
			parts = append(parts, word)
		case 2:
			number, err := rpc.GenerateNumber()
			if err != nil {
				return "", fmt.Errorf("error while calling service: %v", err)
			}
			parts = append(parts, number)
		case 3:
			word, err := rpc.GetWord()
			if err != nil {
				return "", fmt.Errorf("error while calling service: %v", err)
			}
			parts = append(parts, word)
		}
		i = rand.Intn(4)
	}

	return strings.Join(parts, "-"), nil
}
