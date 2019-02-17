package main

import (
	"log"

	"github.com/yb172/experiments/kube/gateway/cfg"
)

func main() {
	if err := cfg.InitConfig(); err != nil {
		log.Fatalf("Unable to init config: %s", err)
	}
}
