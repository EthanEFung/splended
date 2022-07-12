package models

import "errors"

var ErrNotStarted = errors.New("game not started")
var ErrInsufficientPlayers = errors.New("not enough players")

type GameStatePending struct {
	Game *Game
}

func NewGameStatePending(game *Game) GameStatePending {
	return GameStatePending{game}
}

func (s GameStatePending) String() string {
	return "pending"
}

func (s GameStatePending) StartGame() error {
	if len(s.Game.Players) < 2 {
		return ErrInsufficientPlayers
	}
	first := s.Game.Players[0]
	s.Game.SetState(NewGameStatePlaying(s.Game, first))
	return nil
}

func (s GameStatePending) AddPlayer(p *Player) error {
	s.Game.Players = append(s.Game.Players, p)
	return nil
}

func (s GameStatePending) RemovePlayer(p *Player) error {
	if len(s.Game.Players) <= 1 {
		s.Game.Players = []*Player{}
	}
	j := len(s.Game.Players) - 1
	for i, player := range s.Game.Players {
		if p == player {
			s.Game.Players[i] = s.Game.Players[j]
		  s.Game.Players = s.Game.Players[:j]
		}
	}
	return nil
}

func (s GameStatePending) Take3(a, b, c *Token) error {
	return ErrNotStarted
}

func (s GameStatePending) Take2(token *Token) error {
  return ErrNotStarted
}

func (s GameStatePending) Reserve(card *Card) error {
	return ErrNotStarted
}

func (s GameStatePending) Purchase(card *Card) error {
	return ErrNotStarted
}

func (s GameStatePending) EndTurn() error {
	return ErrNotStarted
}