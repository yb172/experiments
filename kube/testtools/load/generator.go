package load

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/yb172/experiments/kube/testtools/cfg"
)

var rate int

// GenerateLoad generates load on our services
func GenerateLoad(attachKeyboard bool) error {
	rate = cfg.Conf.Default.RPS
	exit := make(chan interface{})
	log.Printf("Starting to generate load @%v RPS", rate)
	go generator(exit)
	if attachKeyboard {
		readKeyboard(exit)
	} else {
		<-exit
	}
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()
	address := fmt.Sprintf("http://%s:%s", cfg.Conf.Gen.Gateway.Service.Host, cfg.Conf.Gen.Gateway.Service.Port)
	resp, err := http.Get(address)
	if err != nil {
		log.Printf("error while making request: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Status is not OK: %v", resp.StatusCode)
	}
}
