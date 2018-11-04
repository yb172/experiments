package thegame

import "github.com/looplab/fsm"

const (
	cratedState     = "created"
	waitingState    = "waiting for second player"
	xMoveState      = "waiting for X move"
	xMoveCheckState = "checking X move"
	oMoveState      = "waiting for O move"
	oMoveCheckState = "checking O move"
	xWonState       = "X won"
	oWonState       = "O won"
	drawState       = "draw"

	xJoinsEvent      = "X joins"
	oJoinsEvent      = "O joins"
	xMovesEvent      = "X moves"
	passMoveToOEvent = "pass move to O"
	oMovesEvent      = "O moves"
	passMoveToXEvent = "pass move to X"
	xWinsEvent       = "X wins"
	oWinsEvent       = "O wins"
	drawEvent        = "draw"
)

// GameMatch is an instance of the game
type GameMatch struct {
	fsm *fsm.FSM
}

func newGame() *GameMatch {
	machine := fsm.NewFSM(
		cratedState,
		[]fsm.EventDesc{
			{Name: xJoinsEvent, Src: []string{cratedState}, Dst: waitingState},
			{Name: oJoinsEvent, Src: []string{waitingState}, Dst: xMoveState},
			{Name: xMovesEvent, Src: []string{xMoveState}, Dst: xMoveCheckState},
			{Name: xWinsEvent, Src: []string{xMoveCheckState}, Dst: xWonState},
			{Name: passMoveToOEvent, Src: []string{xMoveCheckState}, Dst: oMoveState},
			{Name: oMovesEvent, Src: []string{oMoveState}, Dst: oMoveCheckState},
			{Name: oWinsEvent, Src: []string{oMoveCheckState}, Dst: oWonState},
			{Name: passMoveToXEvent, Src: []string{oMoveCheckState}, Dst: xMoveState},
			{Name: drawEvent, Src: []string{xMoveCheckState, oMoveCheckState}, Dst: drawState},
		},
		map[string]fsm.Callback{},
	)
	return &GameMatch{
		fsm: machine,
	}
}
