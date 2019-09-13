package FastExp

import (
	"math/big"
)

func NaivePowWithMod(x, y, m int64) int64 {
	X := big.NewInt(x)
	Y := big.NewInt(y)
	M := big.NewInt(m)

	return big.NewInt(0).Exp(X, Y, M).Int64()
}

// Modular exponentiation
// x ** y % |m|
func FastExp(x, y, m int64) int64 {
	if y < 0 {
		if x < 1 {
			return 0
		}

		if m < 0 {
			return 1
		}

		return 1
	}

	var indexesOfBits []int

	for i := 0; i < 65; i++ {
		if ((y >> uint(i)) & 1) == 1 {
			indexesOfBits = append(indexesOfBits, i)
		}
	}

	var prevRes []int64

	prevRes = append(prevRes, x % m) // i == 0

	for i := 1; i <= indexesOfBits[len(indexesOfBits) - 1]; i++ {
		prevRes = append(prevRes, ((prevRes[i - 1] % m) * (prevRes[i - 1] % m)) % m)
	}

	var result int64 = 1

	for i := 0; i < len(indexesOfBits); i++ {
		result *= prevRes[indexesOfBits[i]]
		result %= m
	}

	return result % m
}

func SmallFastExp(x, y, m uint64) (result uint64) {
	result = 1

	for ; y != 0; y >>= 1 {
		if y & 1 == 1 {
			result *= x % m
		}
		result %= m
		x = (x * x) % m
	}

	return
}