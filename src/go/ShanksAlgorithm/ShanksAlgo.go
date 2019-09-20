package ShanksAlgorithm

import (
	"cryptocrouse/src/go/FastExp"
	"math"
)

// (a ** x) mod p = y === lSet = rSet
// a ** (i * m) = y * (a ** j)
func ShanksAlgo(a, p, y uint64) (x uint64) {
	if p <= y {
		return 0
	}

	var n = uint64(math.Sqrt(float64(p))) + 1

	values := make(map[uint64]uint64)

	values[1] = 0

	var i uint64
	for i = 1; i < n; i++ {
		values[FastExp.SmallFastExp(a, i * n, p)] = i
	}

	var j uint64
	for j = 0; j <= n; j++ {
		current := (FastExp.SmallFastExp(a, j, p) * y) % p
		if i, ok := values[current]; ok {
			x = i * n - j
		}
	}

	return
}
