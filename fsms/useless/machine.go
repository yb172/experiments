package useless

import (
	"fmt"
	"log"

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
			"before_event": m.log,
			"leave_state":  m.log,
			"enter_state":  m.log,
			"after_event":  m.log,
			"leave_off":    func(_ *fsm.Event) { m.onTurnOn() },
		},
	)

	return &m
}

func (m *MostUselessMachine) log(e *fsm.Event) {
	fmt.Printf("Event: %v\n", e)
}

func (m *MostUselessMachine) onTurnOn() {
	fmt.Println("onTurnOn")
	if err := m.FSM.Event("turn-off"); err != nil {
		log.Println(err)
	}
}
