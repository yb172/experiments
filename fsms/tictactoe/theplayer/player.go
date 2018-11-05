package theplayer

import (
	fmt "fmt"

	"github.com/yb172/experiments/fsms/tictactoe/thegame"
)

// Player provides methods to play the game
type Player struct {
	match *thegame.GameMatch
	label thegame.Label
}

// Connect would connect to existing game or start new game
func Connect() (*Player, error) {
	game, err := thegame.Get()
	if err != nil {
		game, err = thegame.Start()
		if err != nil {
			return nil, fmt.Errorf("unable to get game in progress or start a new one: %v", err)
		}
	}
	label, err := game.Join()
	if err != nil {
		return nil, fmt.Errorf("unable to join the game: %v", err)
	}
	return &Player{
		match: game,
		label: label,
	}, nil
}

// Move makes the move
func (p *Player) Move(row, col int) error {
	if err := p.match.Move(row, col, p.label); err != nil {
		return fmt.Errorf("error while making move to (%v,%v) for %s: %v", row, col, string(p.label), err)
	}
	return nil
}
