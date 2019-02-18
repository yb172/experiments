package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yb172/experiments/kube/gateway/cfg"
	"github.com/yb172/experiments/kube/gateway/gen"
)

func main() {
	if err := cfg.InitConfig(); err != nil {
		log.Fatalf("Unable to init config: %s", err)
	}

	address := fmt.Sprintf(":%v", cfg.Conf.Own.Port)
	http.HandleFunc("/", serve)
	log.Printf("Start serving on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func serve(w http.ResponseWriter, r *http.Request) {
	seq, err := gen.GenerateSeq()
	if err != nil {
		log.Printf("Error happened while generating: %v", err)
		http.Error(w, fmt.Sprintf("Error while generating seq: %v", err), http.StatusInternalServerError)
	}
	_, err = fmt.Fprintf(w, seq)
	if err != nil {
		log.Printf("Error while writing response: %v", err)
	}
}
