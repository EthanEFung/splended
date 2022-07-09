package models

import "fmt"

type GameState int64

const (
	pending GameState = iota
	playing
	played
)

type Game struct {
	State GameState
	Players []*Player
	Current *Player
	Winner *Player
}

func NewGame() *Game {
	game := new(Game)
	game.State = pending

	return game
}

func (g *Game) String() string {
	if g.State == pending {
		return "Waiting to start"
	}
	if g.State == played {
		return fmt.Sprintf("Winner: %v", *g.Winner)
	}
	return "Game is underway"
}