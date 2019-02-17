package gen

import (
	"fmt"
	"log"
	"net"

	"github.com/yb172/experiments/kube/number/cfg"
	"github.com/yb172/experiments/kube/number/proto/wordgen"
	"google.golang.org/grpc"
)

// StartServer starts the server
func StartServer() error {
	s := grpc.NewServer()
	svc := Service{}
	wordgen.RegisterNumberGeneratorServer(s, &svc)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Conf.Own.Port))
	if err != nil {
		return fmt.Errorf("failed to listen: %s", err)
	}
	log.Printf("Start listening on %v", cfg.Conf.Own.Port)
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to start server: %s", err)
	}

	return nil
}
