package models

/*
	GamesState in which a player has too many tokens and must
	return tokens
*/
type GameStateTrading struct {
	Game *Game
	Player *Player
}