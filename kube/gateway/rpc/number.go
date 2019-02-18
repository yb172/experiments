package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/yb172/experiments/kube/gateway/cfg"
	"github.com/yb172/experiments/kube/gateway/proto/wordgen"
	"google.golang.org/grpc"
)

// GenerateNumber requests number
func GenerateNumber() (string, error) {
	address := fmt.Sprintf("%s:%v", cfg.Conf.Gen.Number.Service.Host, cfg.Conf.Gen.Number.Service.Port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return "", fmt.Errorf("error while connecting: %v", err)
	}
	defer conn.Close()
	c := wordgen.NewNumberGeneratorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := c.GenerateNumber(ctx, &wordgen.GenerateNumberReq{Min: 0, Max: 100})
	if err != nil {
		return "", fmt.Errorf("error while making request: %v", err)
	}
	return fmt.Sprintf("%v", resp.Number), nil
}
