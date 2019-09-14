package Diffie_Hellman

import (
	"math"
	"math/rand"
)

func IsPrime(x uint64) bool {
	if x <= 1 {
		return false
	}

	b := uint64(math.Pow(float64(x), 0.5))

	var i uint64
	for i = 2; i <= b; i++ {
		if (x % i) == 0 {
			return false
		}
	}

	return true
}

func generatePrimeNumber() (result uint64) {
	result = 0
	for ; IsPrime(result); result = rand.Uint64() {
	}
	return
}
