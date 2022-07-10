package models

import (
	"errors"
	"log"
)

type Game struct {
	State GameState
	Decks [][]*Card
	Board [][]*Card
	Tokens map[Gem][]*Token // groups of non gold gems
	Jokers []*Token // list of gold
	Nobles []*Noble
	Players []*Player
	Current *Player
	Winner *Player
}

func NewGame() *Game {
	game := new(Game)
	game.State = PendingState{game}
	return game
}

// really this should be restricted to the Pending state
func (g *Game) AddPlayer(p *Player) error {
	if g.CurrentState() != "pending" {
    return errors.New("cannot add player at this time")
	}
	g.Players = append(g.Players, p)
	return nil
}
// really this should be restricted to the Pending state
func (g *Game) RemovePlayer(p *Player) error{
	if g.CurrentState() != "pending" {
		return errors.New("cannot remove player at this time")
	}
	if len(g.Players) <= 1 {
		g.Players = []*Player{}
	}
	j := len(g.Players) - 1
	for i, player := range g.Players {
		if p == player {
			g.Players[i] = g.Players[j]
		  g.Players = g.Players[:j]
		}
	}
	return nil
}

func (g *Game) CurrentState() string {
	return g.State.String()
}

func (g *Game) Next() (GameState, error) {
	state, err :=  g.State.Next()
	if err != nil {
		log.Printf("state error after play: %v\n", err)
	}
	return state, err
}

type GameState interface {
	String() string
	Next() (GameState, error)
}

type PendingState struct {
	Game *Game
}
func (pending PendingState) String() string {
	return "pending"
}
func (pending PendingState) Next() (GameState, error) {
	if len(pending.Game.Players) < 2 {
		return pending, errors.New("game requires two players to start")
	}
	first := pending.Game.Players[0]
	playing := PlayingState{pending.Game, first}
	pending.Game.State = playing
	pending.Game.Current = first
	return playing, nil
}

// in the playing state, a user should be able to
// take a turn by manipulating the game
type PlayingState struct {
	Game *Game
	Player *Player
}
func (playing PlayingState) String() string {
	return "playing"
}
func (playing PlayingState) Next() (GameState, error) {
	return playing, nil
}
func (playing PlayingState) Take3(a, b, c *Token) error {
	if a.Value == b.Value && a != nil || b.Value == c.Value && b != nil || c.Value == a.Value && c != nil {
		return ErrThreeTokenRule
	}

	return nil
}
func (playing PlayingState) Take2(token *Token) error {
  if len(playing.Game.Tokens[token.Value]) < 4 {
		return ErrTwoTokenRule
	}
	return nil
}
func (playing PlayingState) Reserve(card *Card) error {
	if len(playing.Player.Reserved) > 2 {
		return ErrReservationRule
	}
	for y := range playing.Game.Board {
		for  x := range playing.Game.Board[y] {

			if card == playing.Game.Board[y][x] {
				
				playing.Player.Reserved = append(playing.Player.Reserved, card)

				l := len(playing.Game.Decks[y])

				if l == 0 {
					return nil
				}

				top := playing.Game.Decks[y][l-1]

				playing.Game.Board[y][x] = top

				playing.Game.Decks[y] = playing.Game.Decks[y][:l-1]

				return nil
			}
		}
	}
	
	return errors.New("could not find the card to reserve")
}
func (playing PlayingState) Purchase(card *Card) error {
	// check for payment
	grossPrice := card.Price
	player := playing.Player
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

	for y := range playing.Game.Board {
		for _, c := range playing.Game.Board[y] {
			if c == card {
				// take payment
				
				
				


				// take card
				gem := card.Bonus
				playing.Player.Cards[gem] = append(playing.Player.Cards[gem], card)

			}
		}
	}

	return errors.New("could not purchase card")
}

type PlayedState struct {
	Game *Game
}