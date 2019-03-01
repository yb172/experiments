package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/yb172/experiments/kube/gateway/cfg"
	"github.com/yb172/experiments/kube/gateway/proto/wordgen"
	"google.golang.org/grpc"
)

// GenerateWord requests word from word generator
func GenerateWord() (string, error) {
	address := fmt.Sprintf("%s:%v", cfg.Conf.Word.Service.Host, cfg.Conf.Word.Service.Port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return "", fmt.Errorf("error while connecting: %v", err)
	}
	defer conn.Close()
	c := wordgen.NewWordGeneratorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := c.GenerateWord(ctx, &wordgen.GenerateWordReq{})
	if err != nil {
		return "", fmt.Errorf("error while making request: %v", err)
	}
	return resp.Word, nil
}
