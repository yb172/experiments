package gen

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/yb172/experiments/kube/internets/proto/wordgen"
)

func TestSmoke(t *testing.T) {
	g := NewGomegaWithT(t)

	svc := Service{}
	resp, err := svc.GetWord(context.Background(), &wordgen.GetWordReq{})
	g.Expect(err).To(BeNil())
	word := resp.Word
	g.Expect(len(word)).Should(BeNumerically(">=", 1))
}

func TestGeneration(t *testing.T) {
	g := NewGomegaWithT(t)

	svc := Service{}
	resp, err := svc.GetWord(context.Background(), &wordgen.GetWordReq{})
	g.Expect(err).To(BeNil())
	word1 := resp.Word
	resp, err = svc.GetWord(context.Background(), &wordgen.GetWordReq{})
	g.Expect(err).To(BeNil())
	word2 := resp.Word

	g.Expect(word1).ToNot(Equal(word2))
}
