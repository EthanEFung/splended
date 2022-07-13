package models

import "testing"

func TestGameStatePlayingRestrictions(t *testing.T) {
	game := NewGame()
	a, b := &Player{}, &Player{}
	game.AddPlayer(a)

	state := NewGameStatePlaying(game, a)
	if err := state.StartGame(); err == nil {
		t.Fatalf("game in playing state should not allow StartGame to be rerun")
	}
	if err := state.AddPlayer(b); err == nil {
		t.Fatalf("game in playing state should not allow new players to join")
	}
	if err := state.RemovePlayer(a); err == nil {
		t.Fatalf("game in playing state should not allow players to be removed")
	}
}