package models

type Game struct {

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

	game.CurrentState = NewGameStatePending(game)
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

func (g *Game) Take3(a, b, c *Token) error {
	return g.CurrentState.Take3(a,b,c)
}

func (g *Game) Take2(token *Token) error {
	return g.CurrentState.Take2(token)
}

func (g *Game) Reserve(card *Card) error {
	return g.CurrentState.Reserve(card)
}

func (g *Game) Purchase(card *Card) error {
	return g.CurrentState.Purchase(card)
}

func (g *Game) EndTurn() error {
	return g.CurrentState.EndTurn()
}
