package gen

import (
	"context"

	"github.com/yb172/experiments/kube/internets/proto/wordgen"
)

// Server provides generation services
type Server struct {
}

// GetWord generates word
func (s *Server) GetWord(ctx context.Context, in *wordgen.GetWordReq) (*wordgen.GetWordResp, error) {
	return nil, nil
}
