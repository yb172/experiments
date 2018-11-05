package thegame

import (
	fmt "fmt"
	"strings"
)

// board size
const size = 3

// Label is a player's label: X or O
type Label rune

const (
	// N is a label for empty space
	N Label = ' '
	// X is a label for first player
	X Label = 'x'
	// O is a label for second player
	O Label = 'o'
)

// board represents game board
type board struct {
	cells [][]Label
}

// createBoard creates new board
func createBoard() *board {
	b := &board{}
	for c := 0; c < size; c++ {
		var col []Label
		for c := 0; c < size; c++ {
			col = append(col, N)
		}
		b.cells = append(b.cells, col)
	}
	return b
}

// move records move of player Label to given cell
func (b *board) move(col, row int, label Label) error {
	if label != X && label != O {
		return fmt.Errorf("Label %s is not supported (only %s and %s are)", string(label), string(X), string(O))
	}
	if row < 0 || row >= size {
		return fmt.Errorf("row value (%v) is out of range [0,%v)", row, size)
	}
	if col < 0 || col >= size {
		return fmt.Errorf("col value (%v) is out of range [0,%v)", col, size)
	}
	if b.cells[col][row] != N {
		return fmt.Errorf("cell (%v,%v) have already been used", row, col)
	}
	b.cells[col][row] = label
	return nil
}

// checkWin checks if given player won with his last move to given cell
func (b *board) checkWin() bool {
	// Check rows
	for r := 0; r < size; r++ {
		if b.cells[0][r] != N &&
			b.cells[0][r] == b.cells[1][r] &&
			b.cells[1][r] == b.cells[2][r] {
			return true
		}
	}

	// Check cols
	for c := 0; c < size; c++ {
		if b.cells[c][0] != N &&
			b.cells[c][0] == b.cells[c][1] &&
			b.cells[c][1] == b.cells[c][2] {
			return true
		}
	}

	// Check diagonals
	if b.cells[0][0] != N &&
		b.cells[0][0] == b.cells[1][1] &&
		b.cells[1][1] == b.cells[2][2] {
		return true
	}
	if b.cells[0][2] != N &&
		b.cells[0][2] == b.cells[1][1] &&
		b.cells[1][1] == b.cells[2][0] {
		return true
	}

	return false
}

// checkDraw check if there are any free cells
func (b *board) checkDraw() bool {
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if b.cells[c][r] == N {
				return false
			}
		}
	}
	return true
}

func (b *board) String() string {
	var rows []string
	for r := 0; r < size; r++ {
		var row []string
		for c := 0; c < size; c++ {
			row = append(row, fmt.Sprintf("%s", string(b.cells[c][r])))
		}
		rows = append(rows, strings.Join(row, ""))
	}
	return fmt.Sprintf("%s\n", strings.Join(rows, "\n"))
}
