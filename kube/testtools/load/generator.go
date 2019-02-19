package load

import (
	"log"
	"net/http"
	"time"

	"github.com/yb172/experiments/kube/testtools/cfg"
)

var rate int

// GenerateLoad generates load on our services
func GenerateLoad() error {
	exit := make(chan interface{})
	go generator(exit)
	readKeyboard(exit)
	return nil
}

func generator(exit <-chan interface{}) error {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			for i := 0; i < rate; i++ {
				go makeRequest()
			}
		case <-exit:
			return nil
		}
	}
}

func makeRequest() {
	resp, err := http.Get(cfg.Conf.Service.Address)
	if err != nil {
		log.Printf("error while making request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("Status is not OK: %v", resp.StatusCode)
	}
}
