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

func IsPrimeGoogle(x uint64) bool {
	if x <= 1 {
		return false;
	}

	if x % 2 == 0 {
		return x == 2
	}

	var i uint64
	for i = 3;  ; i += 2 {
		if i > x / i {
			break
		}

		if x % i == 0 {
			return false
		}
	}

	return true
}

func GeneratePrimeNumber() (result uint64) {
	result = rand.Uint64() % MaxBound
	for !IsPrime(result) {
		result = rand.Uint64() % MaxBound
		if result < 2 {
			continue
		}
	}
	return
}

func GeneratePrimeNumberWithoutBounds() (result uint64) {
	result = rand.Uint64()
	for !IsPrime(result) {
		result = rand.Uint64()
		if result < 2 {
			continue
		}
	}
	return
}

func GeneratePrimeNumberWithBounds(bound uint64) (result uint64) {
	for {
		result = rand.Uint64() % bound
		if result < 2 {
			continue
		}
		if !IsPrime(result) {
			break
		}
	}
	return
}
