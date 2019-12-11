package main

import "math/big"

var (
	MIN_P = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(16), nil)
	MAX_P = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(32), nil)
)
