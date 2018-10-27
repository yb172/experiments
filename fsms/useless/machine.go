package useless

import (
	"fmt"

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
			"enter_state": func(e *fsm.Event) { m.enterState(e) },
		},
	)

	return &m
}

func (m *MostUselessMachine) enterState(e *fsm.Event) {
	fmt.Printf("Moving from %s to %s\n", e.Src, e.Dst)
}
