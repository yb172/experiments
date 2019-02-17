package main

import (
	"fmt"
	"log"
	"net"

	"github.com/yb172/experiments/kube/word/cfg"
	"github.com/yb172/experiments/kube/word/gen"
	"github.com/yb172/experiments/kube/word/proto/wordgen"
	"google.golang.org/grpc"
)

func main() {
	if err := cfg.InitConfig(); err != nil {
		log.Fatalf("Unable to init config: %s", err)
	}

	s := grpc.NewServer()
	svc := gen.Server{}
	wordgen.RegisterWordGeneratorServer(s, &svc)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Conf.Own.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}
	log.Printf("Start listening on %v", cfg.Conf.Own.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
