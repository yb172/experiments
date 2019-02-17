package gen

import (
	"context"

	"github.com/yb172/experiments/kube/word/proto/wordgen"
)

// Server provides generation services
type Server struct {
}

// GenerateWord generates word
func (s *Server) GenerateWord(ctx context.Context, in *wordgen.GenerateWordReq) (*wordgen.GenerateWordResp, error) {
	return nil, nil
}
