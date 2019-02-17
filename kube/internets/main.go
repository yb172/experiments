package main

import (
	"log"

	"github.com/yb172/experiments/kube/internets/cfg"
	"github.com/yb172/experiments/kube/internets/gen"
)

func main() {
	if err := cfg.InitConfig(); err != nil {
		log.Fatalf("Unable to init config: %s", err)
	}

	if err := gen.StartServer(); err != nil {
		log.Fatalf("Unable to start server: %v", err)
	}
}
