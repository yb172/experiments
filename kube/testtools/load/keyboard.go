package load

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yb172/experiments/kube/testtools/cfg"
)

const rpsIncrement = 2

func readKeyboard(exit chan<- interface{}) error {
	rate = cfg.Conf.Default.RPS

	fmt.Printf("Generating load to service at a rate of %v rps\n", rate)
	fmt.Printf("Please type \"up\" to increase rps, \"down\" to decrease or Ctrl+C to exit\n")

	read()
	exit <- true
	return nil
}

func read() {
	// TODO: Change to react to keypress, not command
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		switch in.Text() {
		case "up":
			upRps()
		case "down":
			downRps()
		case "exit":
			return
		default:
			fmt.Printf("Command not recognized\n")
		}
	}
}

func upRps() {
	rate += rpsIncrement
	fmt.Printf("RPS increased to %v\n", rate)
}

func downRps() {
	if rate > 0 {
		rate -= rpsIncrement
	}
	if rate < 0 {
		rate = 0
	}
	fmt.Printf("RPS decreased to %v\n", rate)
}
