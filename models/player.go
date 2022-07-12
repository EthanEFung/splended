package models

type Player struct {
	Name string
	Cards map[Gem][]*Card
	Tokens map[Gem][]*Token
	Reserved []*Card
}

func NewPlayer() *Player {
	return new(Player)
}

func (p Player) Points() {

}