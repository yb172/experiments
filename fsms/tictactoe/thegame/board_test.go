package thegame

import "testing"

func TestBoardXWinsInCol(t *testing.T) {
	b := createBoard()
	b.move(0, 0, X)
	if b.checkWin() {
		t.Errorf("win is not expected")
	}
	b.move(0, 1, X)
	if b.checkWin() {
		t.Errorf("win is not expected")
	}
	b.move(0, 2, X)
	if b.checkDraw() {
		t.Errorf("draw is not expected")
	}
	if !b.checkWin() {
		t.Errorf("win for is expected")
	}
}

func TestBoardOWinsInRow(t *testing.T) {
	b := createBoard()
	b.move(0, 0, O)
	if b.checkWin() {
		t.Errorf("win is not expected")
	}
	b.move(1, 0, O)
	if b.checkWin() {
		t.Errorf("win for %s is not expected", string(O))
	}
	b.move(2, 0, O)
	if b.checkDraw() {
		t.Errorf("draw is not expected")
	}
	if !b.checkWin() {
		t.Errorf("win is expected")
	}
}

func TestBoardXWinsInDiagonal(t *testing.T) {
	b := createBoard()
	b.move(0, 0, X)
	if b.checkWin() {
		t.Errorf("win is not expected")
	}
	b.move(1, 1, X)
	if b.checkWin() {
		t.Errorf("win is not expected")
	}
	b.move(2, 2, X)
	if b.checkDraw() {
		t.Errorf("draw is not expected")
	}
	if !b.checkWin() {
		t.Errorf("win is expected")
	}
}

func TestBoardNotWinDiagonal(t *testing.T) {
	b := createBoard()
	b.move(0, 0, X)
	if b.checkWin() {
		t.Errorf("win is not expected")
	}
	b.move(1, 1, X)
	if b.checkWin() {
		t.Errorf("win is not expected")
	}
	b.move(2, 2, O)
	if b.checkDraw() {
		t.Errorf("draw is not expected")
	}
	if b.checkWin() {
		t.Errorf("win is not expected")
	}
}

func TestBoardDraw(t *testing.T) {
	b := createBoard()
	b.move(0, 0, X)
	b.move(0, 1, O)
	b.move(0, 2, X)
	b.move(1, 0, X)
	b.move(1, 1, O)
	b.move(1, 2, X)
	b.move(2, 0, O)
	b.move(2, 1, X)
	b.move(2, 2, O)
	if !b.checkDraw() {
		t.Errorf("draw is expected")
	}
}

func TestBoardMoveToSameCell(t *testing.T) {
	b := createBoard()
	b.move(0, 0, X)
	err := b.move(0, 0, O)
	if err == nil {
		t.Errorf("error is expected")
	}
}

func TestBoardInvalidMoveParams(t *testing.T) {
	b := createBoard()
	err := b.move(-1, 0, X)
	if err == nil {
		t.Errorf("error is expected")
	}
	err = b.move(0, 3, X)
	if err == nil {
		t.Errorf("error is expected")
	}
	err = b.move(0, 2, Label('s'))
	if err == nil {
		t.Errorf("error is expected")
	}
}
