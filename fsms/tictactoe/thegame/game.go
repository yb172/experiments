package thegame

import (
	fmt "fmt"

	"github.com/looplab/fsm"
)

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
	fsm   *fsm.FSM
	board *board
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
		map[string]fsm.Callback{
			// Transition from X move is async since move might be illegal
			// (then transition would be cancelled)
			fmt.Sprintf("leave_%s", xMoveState): validateMove,
			// Transition from O move is async since move might be illegal
			// (then transition would be cancelled)
			fmt.Sprintf("leave_%s", oMoveState): validateMove,
			"enter_state": func(e *fsm.Event) {
				fmt.Printf("Move from %s to %s\n", e.Src, e.Dst)
			},
		},
	)
	return &GameMatch{
		fsm:   machine,
		board: createBoard(),
	}
}

func validateMove(e *fsm.Event) {
	board := e.Args[0].(*board)
	col := e.Args[1].(int)
	row := e.Args[2].(int)
	label := e.Args[3].(Label)
	err := board.move(col, row, label)
	if err != nil {
		e.Cancel(fmt.Errorf("unable to perform move to (%v,%v) for %s: %v", row, col, string(label), err))
	}
}

// Join joins the game or returns an error if it's not possible
func (g *GameMatch) Join() (Label, error) {
	var label Label
	switch {
	case g.fsm.Can(xJoinsEvent):
		label = X
	case g.fsm.Can(oJoinsEvent):
		label = O
	default:
		return N, fmt.Errorf("unable to join the game in current state: %s", g.fsm.Current())
	}
	err := g.fsm.Event(joinFor(label))
	if err != nil {
		return N, fmt.Errorf("not able to join the game as %s: %v", string(label), err)
	}
	return label, nil
}

// Move makes move for player
func (g *GameMatch) Move(col, row int, label Label) error {
	fmt.Println(g.board)
	if err := g.fsm.Event(moveFor(label), g.board, col, row, label); err != nil {
		return fmt.Errorf("unable to perform move for %s: %v", string(label), err)
	}
	switch {
	case g.board.checkWin():
		if err := g.fsm.Event(winFor(label)); err != nil {
			return fmt.Errorf("error while transitioning to win state for %s: %v", string(label), err)
		}
	case g.board.checkDraw():
		if err := g.fsm.Event(drawEvent); err != nil {
			return fmt.Errorf("error while transitioning to draw state: %v", err)
		}
	default:
		if err := g.fsm.Event(passFor(label)); err != nil {
			return fmt.Errorf("error while passing move to next player for %s: %v", string(label), err)
		}
	}
	return nil
}
