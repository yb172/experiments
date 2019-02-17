package gen

import (
	"fmt"
	"log"
	"net"

	"github.com/yb172/experiments/kube/word/cfg"
	"github.com/yb172/experiments/kube/word/proto/wordgen"
	"google.golang.org/grpc"
)

// StartServer starts the server
func StartServer() error {
	s := grpc.NewServer()
	svc := Service{}
	wordgen.RegisterWordGeneratorServer(s, &svc)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Conf.Own.Port))
	if err != nil {
		return fmt.Errorf("Failed to listen: %s", err)
	}
	log.Printf("Start listening on %v", cfg.Conf.Own.Port)
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("Failed to start server: %s", err)
	}

	return nil
}
