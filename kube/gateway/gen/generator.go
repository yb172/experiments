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
	rand.Seed(time.Now().UTC().UnixNano())

	var parts []string
	for {
		i := rand.Intn(4)
		switch i {
		case 0:
			return strings.Join(parts, "-"), nil
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
	}
}
