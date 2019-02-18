package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/yb172/experiments/kube/gateway/cfg"
	"github.com/yb172/experiments/kube/gateway/proto/wordgen"
	"google.golang.org/grpc"
)

// GetWord requests word from internets
func GetWord() (string, error) {
	address := fmt.Sprintf("%s:%v", cfg.Conf.Gen.Internets.Service.Host, cfg.Conf.Gen.Internets.Service.Port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return "", fmt.Errorf("error while connecting: %v", err)
	}
	defer conn.Close()
	c := wordgen.NewInternetsWordGeneratorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := c.GetWord(ctx, &wordgen.GetWordReq{})
	if err != nil {
		return "", fmt.Errorf("error while making request: %v", err)
	}
	return resp.Word, nil
}
