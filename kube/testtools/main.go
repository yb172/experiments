package main

import (
	"log"

	"github.com/yb172/experiments/kube/testtools/cfg"
	"github.com/yb172/experiments/kube/testtools/cmd"
)

func main() {
	if err := cfg.InitConfig(); err != nil {
		log.Fatalf("Unable to init config: %s", err)
	}

	cmd.Execute()
}
