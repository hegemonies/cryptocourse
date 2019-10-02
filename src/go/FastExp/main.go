package FastExp

import (
	"math/big"
)

func NaivePowWithMod(a, x, p int64) int64 {
	A := big.NewInt(a)
	X := big.NewInt(x)
	P := big.NewInt(p)

	return big.NewInt(0).Exp(A, X, P).Int64()
}

// Modular exponentiation
// a ** x % |p|
func FastExp(a, x, p int64) int64 {
	if x < 0 {
		if a < 1 {
			return 0
		}

		if p < 0 {
			return 1
		}

		return 1
	}

	var indexesOfBits []int

	for i := 0; i < 65; i++ {
		if ((x >> uint(i)) & 1) == 1 {
			indexesOfBits = append(indexesOfBits, i)
		}
	}

	var prevRes []int64

	prevRes = append(prevRes, a % p) // i == 0

	for i := 1; i <= indexesOfBits[len(indexesOfBits) - 1]; i++ {
		prevRes = append(prevRes, ((prevRes[i - 1] % p) * (prevRes[i - 1] % p)) % p)
	}

	var result int64 = 1

	for i := 0; i < len(indexesOfBits); i++ {
		result *= prevRes[indexesOfBits[i]]
		result %= p
	}

	return result % p
}

func SmallFastExp(a, x, p uint64) (result uint64) {
	result = 1

	for ; x != 0; x >>= 1 {
		if x & 1 == 1 {
			result *= a % p
		}
		result %= p
		a = (a * a) % p
	}

	return
}
