package models

type Noble struct {
	/* Prestige is the number added to the players total points */
	Prestige int
	/* Bonuses is a number of cards the required to obtain by gem */
	Bonuses map[Gem]int
}
