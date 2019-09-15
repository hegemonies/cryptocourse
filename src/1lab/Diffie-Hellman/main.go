package Diffie_Hellman

import (
	"math"
)

func IsPrime(x int64) bool {
	if x <= 1 {
		return false
	}

	b := int64(math.Pow(float64(x),0.5))

	var i int64
	for i = 2; i <= b; i++ {
		if (x % i) == 0 {
			return false
		}
	}

	//big.NewInt(0).ProbablyPrime() todo: use for test

	return true
}
