package models

import (
	"testing"
)

func TestAddPlayer(t *testing.T) {
	game := NewGame()

	a, b := &Player{}, &Player{}
	game.AddPlayer(a)
	if len(game.Players) != 1 {
		t.Fatalf("Could not add a player using the AddPlayer method")
	}
	game.AddPlayer(b)
	if len(game.Players) != 2 {
		t.Fatalf("Could not add a second Player using the AddPlayer method")
	}
}

func TestRemovePlayer(t *testing.T) {
	game := NewGame()

	a, b := &Player{}, &Player{}
	game.AddPlayer(a)
	game.AddPlayer(b)

	game.RemovePlayer(a)
	if game.Players[0] != b {
		t.Fatalf("Attempted to remove a player from the game, but failed")
	}
	game.RemovePlayer(b)
	if len(game.Players) != 0 {
		t.Fatalf("Attempted to remove the last player from the game but failed")
	}
	game.AddPlayer(a)
	game.AddPlayer(b)
	game.RemovePlayer(b)
	if len(game.Players) != 1 {
		t.Fatalf("Attempted to remove the last player of the game, but failed")
	}
}

func TestStartingAGame(t *testing.T) {
	g := NewGame()
	a, b := &Player{}, &Player{}

	g.AddPlayer(a)

	err := g.StartGame()
  if err == nil {
		t.Fatalf("successfully able to start a game with insufficient number of players")
	} 
	if g.StateString() != "pending" {
		t.Fatalf("somehow changed the state when the state should have remained the same")
	}
	g.AddPlayer(b)
	err = g.StartGame()
	if err != nil {
		t.Fatalf("could not start a game with sufficient number of players")
	}
	if g.StateString() == "pending" {
		t.Fatalf("game transition was not made to playing")
	}
}

func TestGamePlay(t *testing.T) {
	g := NewGame()
	a, b := &Player{}, &Player{}
	g.AddPlayer(a)
	g.AddPlayer(b)
	g.StartGame()
	_, ok := g.CurrentState.(GameStatePlaying)
	if !ok {
		t.Fatalf("Could not transition into playing state")
	}

}