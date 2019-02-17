package gen

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/yb172/experiments/kube/number/proto/wordgen"
)

func TestSmoke(t *testing.T) {
	g := NewGomegaWithT(t)

	svc := Service{}
	resp, err := svc.GenerateNumber(context.Background(), &wordgen.GenerateNumberReq{Min: 0, Max: 100})
	g.Expect(err).To(BeNil())
	num := resp.Number
	g.Expect(num).Should(BeNumerically(">=", 0))
}

func TestGeneration(t *testing.T) {
	g := NewGomegaWithT(t)

	svc := Service{}
	resp, err := svc.GenerateNumber(context.Background(), &wordgen.GenerateNumberReq{Min: 0, Max: 100})
	g.Expect(err).To(BeNil())
	num1 := resp.Number
	resp, err = svc.GenerateNumber(context.Background(), &wordgen.GenerateNumberReq{Min: 100, Max: 200})
	g.Expect(err).To(BeNil())
	num2 := resp.Number

	g.Expect(num1).ToNot(Equal(num2))
}
