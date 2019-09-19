package Diffie_Hellman

import (
	"math"
	"math/rand"
)

func IsPrime(x uint64) bool {
	b := uint64(math.Pow(float64(x), 0.5))

	var i uint64
	for i = 2; i <= b; i++ {
		if (x % i) == 0 {
			return false
		}
	}

	return true
}

func GeneratePrimeNumber() (result uint64) {
	result = rand.Uint64() % MaxBound
	for ; !IsPrime(result); result = rand.Uint64() % MaxBound {
		if result == 1 {
			continue
		}
	}
	return
}