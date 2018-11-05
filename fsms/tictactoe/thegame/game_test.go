package thegame

import (
	"strings"
	"testing"
)

func TestGameXWins(t *testing.T) {
	var err error
	g := newGame()
	// X tries to move before joining
	err = g.Move(1, 1, X)
	if err == nil {
		t.Fatalf("error is expected")
	}
	if !strings.Contains(err.Error(), "inappropriate in current state") {
		t.Fatalf("expected error about inapropriate state but got: %v", err)
	}
	_, err = g.Join()
	if err != nil {
		t.Fatalf("error not expected")
	}
	// X tries to move before O joined
	err = g.Move(1, 1, X)
	if err == nil {
		t.Fatalf("error is expected")
	}
	if !strings.Contains(err.Error(), "inappropriate in current state") {
		t.Fatalf("expected error about inapropriate state but got: %v", err)
	}
	_, err = g.Join()
	if err != nil {
		t.Fatalf("error not expected")
	}
	// O tries to move first
	// X tries to move before O joined
	err = g.Move(1, 1, O)
	if err == nil {
		t.Fatalf("error is expected")
	}
	if !strings.Contains(err.Error(), "inappropriate in current state") {
		t.Fatalf("expected error about inapropriate state but got: %v", err)
	}
	// X moves
	// ---
	// -x-
	// ---
	err = g.Move(1, 1, X)
	if err != nil {
		t.Fatalf("error not expected: %v", err)
	}
	// X tries to move again
	err = g.Move(1, 0, X)
	if err == nil {
		t.Fatalf("error is expected")
	}
	if !strings.Contains(err.Error(), "inappropriate in current state") {
		t.Fatalf("expected error about inapropriate state but got: %v", err)
	}
	// O tries to move on cell where X have already moved
	err = g.Move(1, 1, O)
	if err == nil {
		t.Fatalf("error is expected")
	}
	if !strings.Contains(err.Error(), "unable to perform move") {
		t.Fatalf("expected error about unable to move but got: %v", err)
	}
	// O moves
	// o--
	// -x-
	// ---
	err = g.Move(0, 0, O)
	if err != nil {
		t.Fatalf("error not expected: %v", err)
	}
	// X moves
	// ox-
	// -x-
	// ---
	err = g.Move(1, 0, X)
	if err != nil {
		t.Fatalf("error not expected: %v", err)
	}
	// O moves
	// ox-
	// ox-
	// ---
	err = g.Move(0, 1, O)
	if err != nil {
		t.Fatalf("error not expected: %v", err)
	}
	// X moves
	// ox-
	// ox-
	// -x-
	err = g.Move(1, 2, X)
	if err != nil {
		t.Fatalf("error not expected: %v", err)
	}
	// O tries to move
	err = g.Move(0, 2, O)
	if err == nil {
		t.Fatalf("error is expected")
	}
	if !strings.Contains(err.Error(), "inappropriate in current state") {
		t.Fatalf("expected error about inapropriate state but got: %v", err)
	}
	if g.fsm.Current() != xWonState {
		t.Errorf("%q state expected but got %q", xWonState, g.fsm.Current())
	}
}

func TestGameOWins(t *testing.T) {
	var err error
	g := newGame()
	g.Join()
	g.Join()
	err = g.Move(0, 0, X)
	if err != nil {
		t.Fatalf("error not expected: %v", err)
	}
	err = g.Move(1, 0, O)
	if err != nil {
		t.Fatalf("error not expected: %v", err)
	}
	err = g.Move(2, 0, X)
	if err != nil {
		t.Fatalf("error not expected: %v", err)
	}
	// xox
	// ---
	// ---

	err = g.Move(1, 1, O)
	if err != nil {
		t.Fatalf("error not expected: %v", err)
	}
	err = g.Move(0, 1, X)
	if err != nil {
		t.Fatalf("error not expected: %v", err)
	}
	err = g.Move(2, 1, O)
	if err != nil {
		t.Fatalf("error not expected: %v", err)
	}
	// xox
	// xoo
	// ---

	err = g.Move(1, 2, X)
	if err != nil {
		t.Fatalf("error not expected: %v", err)
	}
	err = g.Move(0, 2, O)
	if err != nil {
		t.Fatalf("error not expected: %v", err)
	}
	err = g.Move(2, 2, X)
	if err != nil {
		t.Fatalf("error not expected: %v", err)
	}
	// xox
	// xoo
	// oxx

	// O tries to move
	err = g.Move(2, 0, O)
	if err == nil {
		t.Fatalf("error is expected")
	}
	if !strings.Contains(err.Error(), "inappropriate in current state") {
		t.Fatalf("expected error about inapropriate state but got: %v", err)
	}
	if g.fsm.Current() != drawState {
		t.Errorf("%q state expected but got %q", drawState, g.fsm.Current())
	}
}
