package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/yb172/experiments/kube/gateway/cfg"
	"github.com/yb172/experiments/kube/gateway/rpc"
)

func main() {
	if err := cfg.InitConfig(); err != nil {
		log.Fatalf("Unable to init config: %s", err)
	}

}

func serve(w http.ResponseWriter, r *http.Request) {
	// Seed rand
	rand.Seed(time.Now().UTC().UnixNano())

	i := rand.Intn(4)
	var parts []string
	for i != 0 {
		switch i {
		case 1:
			word, err := rpc.GetWord()
			if err != nil {

			}
			parts = append(parts, word)
		}
	}
}
