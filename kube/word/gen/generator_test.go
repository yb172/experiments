package gen

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/yb172/experiments/kube/word/proto/wordgen"
)

func TestSmoke(t *testing.T) {
	g := NewGomegaWithT(t)

	svc := Service{}
	resp, err := svc.GenerateWord(context.Background(), &wordgen.GenerateWordReq{})
	g.Expect(err).To(BeNil())
	word := resp.Word
	g.Expect(len(word)).Should(BeNumerically(">=", 1))
}

func TestGeneration(t *testing.T) {
	g := NewGomegaWithT(t)

	svc := Service{}
	resp, err := svc.GenerateWord(context.Background(), &wordgen.GenerateWordReq{})
	g.Expect(err).To(BeNil())
	word1 := resp.Word
	resp, err = svc.GenerateWord(context.Background(), &wordgen.GenerateWordReq{})
	g.Expect(err).To(BeNil())
	word2 := resp.Word
	// There is 1 in 2048 chance of generating the same word so we generate one more word
	// There is still a chance of getting the same word, but it's 1 in 2048^2, which is 1 in 4 million
	resp, err = svc.GenerateWord(context.Background(), &wordgen.GenerateWordReq{})
	g.Expect(err).To(BeNil())
	word3 := resp.Word

	g.Expect(word1).Should(Or(Not(Equal(word2)), Not(Equal(word3))))
}
