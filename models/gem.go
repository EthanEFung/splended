package models

/* Gem is a unit of purchase */
type Gem int64

const (
	Emerald Gem = iota
	Diamond
	Sapphire
	Onyx
	Ruby
	Gold
)