package BlindVote

import "math/big"

var (
	MaxP = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(1024), nil)
)
