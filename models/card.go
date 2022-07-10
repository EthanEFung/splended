package models

type Card struct {
	// Level determines which deck the card belongs to
	Level int
	// Prestige is the number added to the players points
	Prestige int
	// Price the purchasing price by gems
	Price map[Gem]int
	// Bonus is the gem used as a discount for subsequent purchases
	Bonus Gem
}

