package gen

import (
	"context"
	"math/rand"
	"time"

	"github.com/yb172/experiments/kube/number/proto/wordgen"
)

// Service provides generation services
type Service struct {
}

// GenerateNumber generates number
func (s *Service) GenerateNumber(context context.Context, req *wordgen.GenerateNumberReq) (*wordgen.GenerateNumberResp, error) {
	rand.Seed(time.Now().UTC().UnixNano())

	num := rand.Int31n(req.Min+req.Max) - req.Min
	return &wordgen.GenerateNumberResp{Number: num}, nil
}
