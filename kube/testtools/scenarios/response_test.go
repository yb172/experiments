package scenarios_test

import (
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"

	"github.com/yb172/experiments/kube/testtools/cfg"
)

var _ = Describe("Response", func() {

	// TODO: Move this to global before
	BeforeEach(func() {
		viper.AddConfigPath("./..")
		if err := cfg.InitConfig(); err != nil {
			panic(err)
		}
	})

	Describe("When generating sequence", func() {
		Context("Sequence is", func() {
			It("Never empty", func() {
				// TODO: Consider making in concurrent
				requests := 20 // It is possible to get all non-empty responses, but probability is low
				for i := 0; i < requests; i++ {
					seq, err := makeRequest()
					if err != nil {
						Fail(fmt.Sprintf("Failed to perform request: %v", err))
					}
					Expect(seq).ToNot(Equal(""))
				}
			})
		})

	})
})

func makeRequest() (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()
	resp, err := http.Get(cfg.Conf.Service.Address)
	if err != nil {
		return "", fmt.Errorf("error while making request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status is not OK: %v", resp.StatusCode)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error while reading response: %v", err)
	}
	return string(bytes), nil
}
