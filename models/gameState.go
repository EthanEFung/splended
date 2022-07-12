package models

type GameState interface {
	String() string
	StartGame() error
	AddPlayer(p *Player) error
	RemovePlayer(p *Player) error
	Take3(a, b, c *Token) error
	Take2(token *Token) error
	Reserve(card *Card) error
	Purchase(card *Card) error
	EndTurn() error 
}
