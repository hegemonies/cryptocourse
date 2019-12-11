package srvr

import "math/big"

const (
	CONN_HOST = "localhost"
	CONN_PORT = "7575"
	CONN_TYPE = "tcp"
)

var (
	MIN_P = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(16), nil)
	MAX_P = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(32), nil)
)
