package ShanksAlgorithm

import (
	. "cryptocrouse/src/1lab/FastExp"
	"math"
)

// (a ** x) mod p = y === lValue = rValue
// a ** (i * m) = y * (a ** j)
func BabyStepGiantStep(a, p, y uint64) (x uint64) {
	m := uint64(math.Pow(float64(p), 0.5))
	k := m

	//lvalue := make([]uint64, 0, k)
	var i uint64
	lvalue := initLvalue(a, p, m, k, &i)

	var j uint64
	rvalue := initRvalue(a, p, y, m, &j)

	x = i * m - j
	return
}

func initLvalue(a, p, m, k uint64, i* uint64) []uint64 {
	lvalue := make([]uint64, 0, k)
	*i = 1
	for ; *i <= k; *i++ {
		lvalue = append(lvalue, SmallFastExp(a, *i * m, p))
	}

	return lvalue
}

func initRvalue(a, p, y, m uint64, j* uint64) []uint64 {
	rvalue := make([]uint64, 0, m)
	rvalue = append(rvalue, y)
	*j = 1
	for ; *j < m; *j++ {
		rvalue = append(rvalue, (rvalue[*j - 1] * a) % p)
		//for // todo: finish it
	}

	return rvalue
}
