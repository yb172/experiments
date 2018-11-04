package thegame

import fmt "fmt"

// Instance of the game.
// As of current implementation only one concurrent game is supported
var match *GameMatch

// Start starts new match
func Start() (*GameMatch, error) {
	if match != nil {
		return nil, fmt.Errorf("game is already in progress")
	}
	return &GameMatch{}, nil
}

// Join joins player to match in progress
func Join() (*GameMatch, error) {
	if match == nil {
		return nil, fmt.Errorf("there is no game in progress")
	}
	return match, nil
}