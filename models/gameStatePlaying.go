package models

import (
	"errors"
	"fmt"
)

type GameStatePlaying struct {
	Game *Game
	Player *Player
}

func NewGameStatePlaying(g *Game, p *Player) *GameStatePlaying {
	return &GameStatePlaying{g, p}
}

func (s GameStatePlaying) String() string {
	return fmt.Sprintf("playing %s's turn", s.Player.Name)
} 

func (s GameStatePlaying) StartGame() error {
	return errors.New("game has been started")
}

func (s GameStatePlaying) AddPlayer(p *Player) error {
	return errors.New("cannot add a player while game is in progress")
}

func (s GameStatePlaying) RemovePlayer(p *Player) error {
	return errors.New("cannot remove a player while a game is in progress")
}

// in the playing state, a user should be able to
// take a turn by manipulating the game
func (s GameStatePlaying) Take3(a, b, c *Token) error {
	if a.Value == b.Value && a != nil || b.Value == c.Value && b != nil || c.Value == a.Value && c != nil {
		return ErrThreeTokenRule
	}
	// TODO: we should be evaluating whether the player has too many tokens
	return nil
}
func (s GameStatePlaying) Take2(token *Token) error {
  if len(s.Game.Tokens[token.Value]) < 4 {
		return ErrTwoTokenRule
	}
	// TODO: we should be evaluating whether the player has too many tokens
	return nil
}
func (s GameStatePlaying) Reserve(card *Card) error {
	if len(s.Player.Reserved) > 2 {
		return ErrReservationRule
	}
	for y := range s.Game.Board {
		for  x := range s.Game.Board[y] {

			if card == s.Game.Board[y][x] {
				
				s.Player.Reserved = append(s.Player.Reserved, card)

				l := len(s.Game.Decks[y])

				if l == 0 {
					return nil
				}

				top := s.Game.Decks[y][l-1]

				s.Game.Board[y][x] = top

				s.Game.Decks[y] = s.Game.Decks[y][:l-1]


				// TODO: we should be evaluating here whether the user 

				return nil
			}
		}
	}
	
	return errors.New("could not find the card to reserve")
}
func (s GameStatePlaying) Purchase(card *Card) error {
	// check for payment
	grossPrice := card.Price
	player := s.Player
	netPrice := make(map[Gem]int)

	for gem, count := range grossPrice {
		netPrice[gem] = count
	}

	for gem, cards := range player.Cards {
		// discounted
		netPrice[gem] -= len(cards)

		if netPrice[gem] < 0 {
			netPrice[gem] = 0
		}

		if netPrice[gem] > len(player.Tokens[gem]) {
			return ErrPurchaseRule
		}

	}

	// now check for card

	for y := range s.Game.Board {
		for _, c := range s.Game.Board[y] {
			if c == card {
				// take payment
				

				// take card
				gem := card.Bonus
				s.Player.Cards[gem] = append(s.Player.Cards[gem], card)

			}
		}
	}

	return errors.New("could not purchase card")
}

func (s GameStatePlaying) EndTurn() error {
	// TODO: here we want to lace in
	// 1. evaluation of the end game (namely that one of the players has met the points
	//	  and the current player is the last)
	current := s.Player
	var next *Player

	for i, player := range s.Game.Players {
		if player == current {
			if i == len(s.Game.Players) - 1 {
				next = s.Game.Players[0]
			} else {
				next = s.Game.Players[i+1]
			}
			break
		}
	}
	if next == nil {
		return errors.New("next player not found")
	}
	playing := NewGameStatePlaying(s.Game, next)
	s.Game.SetState(playing)
	return nil
}