package models

type Game struct {
	Pending GameState
	Playing GameState

	CurrentState GameState

	Decks [][]*Card
	Board [][]*Card
	Tokens map[Gem][]*Token // groups of non gold gems
	Jokers []*Token // list of gold
	Nobles []*Noble
	Players []*Player
}

func NewGame() *Game {
	game := new(Game)
	game.Pending = GameStatePending{game}
	game.Playing = NewGameStatePlaying(game, nil)

	game.CurrentState = game.Pending
	return game
}

func (g *Game) SetState(s GameState) {
	g.CurrentState = s
}

func (g *Game) StateString() string {
	return g.CurrentState.String()
}

func (g *Game) StartGame() error {
	return g.CurrentState.StartGame()
}

func (g *Game) AddPlayer(p *Player) error {
	return g.CurrentState.AddPlayer(p)
}

func (g *Game) RemovePlayer(p *Player) error{
	return g.CurrentState.RemovePlayer(p)
}
