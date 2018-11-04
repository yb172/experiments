package thegame

import "testing"

func TestDrawGame(t *testing.T) {
	m := newGame().fsm
	// X joins
	err := m.Event(xJoinsEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// O joins
	err = m.Event(oJoinsEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// X moves
	err = m.Event(xMovesEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// No win or draw - pass move to O
	err = m.Event(passMoveToOEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// O moves
	err = m.Event(oMovesEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// No win or draw - pass move to X
	err = m.Event(passMoveToXEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// X moves
	err = m.Event(xMovesEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// draw happens
	err = m.Event(drawEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	if m.Current() != drawState {
		t.Fatalf("expected state to be %q", drawState)
	}
}

func TestXWonGame(t *testing.T) {
	m := newGame().fsm
	// X joins
	err := m.Event(xJoinsEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// O joins
	err = m.Event(oJoinsEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// X moves
	err = m.Event(xMovesEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// No win or draw - pass move to O
	err = m.Event(passMoveToOEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// O moves
	err = m.Event(oMovesEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// No win or draw - pass move to X
	err = m.Event(passMoveToXEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// X moves
	err = m.Event(xMovesEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// X wins
	err = m.Event(xWinsEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	if m.Current() != xWonState {
		t.Fatalf("expected state to be %q", xWonState)
	}
}

func TestYWonGame(t *testing.T) {
	m := newGame().fsm
	// X joins
	err := m.Event(xJoinsEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// O joins
	err = m.Event(oJoinsEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// X moves
	err = m.Event(xMovesEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// No win or draw - pass move to O
	err = m.Event(passMoveToOEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// O moves
	err = m.Event(oMovesEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// No win or draw - pass move to X
	err = m.Event(passMoveToXEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// X moves
	err = m.Event(xMovesEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// No win or draw - pass move to O
	err = m.Event(passMoveToOEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// O moves
	err = m.Event(oMovesEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// O wins
	err = m.Event(oWinsEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	if m.Current() != oWonState {
		t.Fatalf("expected state to be %q", oWonState)
	}
}

func TestOTriesToMoveFirst(t *testing.T) {
	m := newGame().fsm
	// X joins
	err := m.Event(xJoinsEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// O joins
	err = m.Event(oJoinsEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// X moves
	err = m.Event(oMovesEvent)
	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestXTriesToMoveTwice(t *testing.T) {
	m := newGame().fsm
	// X joins
	err := m.Event(xJoinsEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// O joins
	err = m.Event(oJoinsEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// X moves
	err = m.Event(xMovesEvent)
	if err != nil {
		t.Fatal("Expected error")
	}
	// No win or draw - pass move to X (again)
	err = m.Event(passMoveToXEvent)
	if err == nil {
		t.Fatal("Expected error")
	}
	// No win or draw - pass move to O
	err = m.Event(passMoveToOEvent)
	if err != nil {
		t.Fatal("Expected no error")
	}
	// No win or draw - pass move to X
	err = m.Event(xMovesEvent)
	if err == nil {
		t.Fatal("Expected error")
	}
}
