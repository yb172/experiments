package useless

import (
	"fmt"
	"log"
	"time"

	"github.com/looplab/fsm"
)

// MostUselessMachine ... ever!
type MostUselessMachine struct {
	FSM *fsm.FSM
}

// Create most useless machine
func Create() *MostUselessMachine {
	m := MostUselessMachine{}
	m.FSM = fsm.NewFSM(
		"off",
		fsm.Events{
			{Name: "turn-on", Src: []string{"off"}, Dst: "on"},
			{Name: "turn-off", Src: []string{"on"}, Dst: "off"},
		},
		fsm.Callbacks{
			"leave_state": onLeave,
			"leave_off":   m.onTurnOn,
		},
	)

	return &m
}

func onLeave(e *fsm.Event) {
	fmt.Printf("\nMoving from state %s to %s\n", e.Src, e.Dst)
}

func (m *MostUselessMachine) onTurnOn(_ *fsm.Event) {
	go func() {
		time.Sleep(2 * time.Second)
		if m.FSM.Current() == "on" {
			fmt.Println("Machine: Not today")
			if err := m.FSM.Event("turn-off"); err != nil {
				log.Printf("Error happened: %v", err)
			}
		}
	}()
}
