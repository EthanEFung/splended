package models

import "errors"

var ErrTenTokenRule = errors.New("Player is only allowed to have ten tokens at the end of their turn")
var ErrThreeTokenRule = errors.New("Player cannot take three coins if two tokens are the same gem")
var ErrTwoTokenRule = errors.New("Player cannot take two tokens of the same color if their are less than 4 to take")
var ErrReservationRule = errors.New("Player cannot reserve more than 3 cards")
var ErrPurchaseRule = errors.New("Player must have the same amount of tokens and bonuses as the cost of the card")