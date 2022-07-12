package models

/*
	GameState in which the game has finished, and a
	winner has been defined
*/
type GameStateEnded struct {
	Game *Game
	Winner *Player
}