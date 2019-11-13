package MentalPoker

import "math/big"

const (
	TWO = iota + 2
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	T // ten
	J // jack
	Q // queen
	K // king
	A // ace
)

var (
	MaxBound = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(16), nil)
	MinP = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(16), nil)
	MaxP =big.NewInt(0).Exp(big.NewInt(2), big.NewInt(32), nil)
)

const (
	MinCountUsers = 2
	MaxCountUsers = 23
	CountCards = 52
)
