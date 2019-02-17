package gen

import (
	"context"

	"github.com/yb172/experiments/kube/number/proto/wordgen"
)

// Server provides generation services
type Server struct {
}

// GenerateNumber generates number
func (s *Server) GenerateNumber(context.Context, *wordgen.GenerateNumberReq) (*wordgen.GenerateNumberResp, error) {
	return nil, nil
}
