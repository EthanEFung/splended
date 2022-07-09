package models

type Card struct {
	Level int
	Prestige int
	Price map[Gem]int
	Bonus Gem
}

