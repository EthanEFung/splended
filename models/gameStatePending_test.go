package models

import "testing"

func TestPendingStateLiberties(t *testing.T) {
	game := NewGame()
  if game.StateString() != "pending" {
	  t.Fatalf("game not initialized with pending state: %v", game.StateString())
	}

	a, b := &Player{}, &Player{}
	game.AddPlayer(a)
	if len(game.Players) != 1 {
		t.Fatalf("could not add a player using the AddPlayer method")
	}
	game.AddPlayer(b)
	if len(game.Players) != 2 {
		t.Fatalf("could not add a second Player using the AddPlayer method")
	}


	game.RemovePlayer(a)
	if game.Players[0] != b {
		t.Fatalf("attempted to remove a player from the game, but failed")
	}
	game.RemovePlayer(b)
	if len(game.Players) != 0 {
		t.Fatalf("attempted to remove the last player from the game but failed")
	}
	game.AddPlayer(a)
	game.AddPlayer(b)
	game.RemovePlayer(b)
	if len(game.Players) != 1 {
		t.Fatalf("attempted to remove the last player of the game, but failed")
	}
}

func TestPendingStateRestrictions(t *testing.T) {
	state := NewGameStatePending(NewGame())
	msg := "restriction was not applied in pending state"
	a, b, c := &Token{Onyx}, &Token{Sapphire}, &Token{Emerald}
	card := &Card{
		Level: 1,
		Prestige: 0,
		Price: make(map[Gem]int),
		Bonus: Onyx,
	}
	if err :=  state.Take3(a,b,c); err == nil {
		t.Fatalf(msg)
	}
	if err :=  state.Take2(a); err == nil {
		t.Fatalf(msg)
	}
	if err :=  state.Reserve(card); err == nil {
		t.Fatalf(msg)
	}
	if err :=  state.Purchase(card); err == nil {
		t.Fatalf(msg)
	}
	if err :=  state.EndTurn(); err == nil {
		t.Fatalf(msg)
	}
}
