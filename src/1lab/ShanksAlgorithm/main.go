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

	var j uint64 = 1
	rvalue := initRvalue(y, m)

	// search a ** (i * m) == (a ** j) * y
	answerIsSearched, i := searchEql(lvalue, rvalue[0])
	for ; answerIsSearched != true && j < m; j++ {
		answerIsSearched, i = searchEql(lvalue, nextRvalueElement(j, a, p, rvalue))
	}

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

func initRvalue(y, m uint64) []uint64 {
	rvalue := make([]uint64, 0, m)
	rvalue = append(rvalue, y)
	return rvalue
}

func nextRvalueElement(j, a, p uint64, rvalue []uint64) uint64 {
	return (rvalue[j - 1] * a) % p
}

func searchEql(lvalue []uint64, lastPrime uint64) (bool, uint64) {
	var i uint64 = uint64(len(lvalue) - 1)
	for ; i > 0; i-- {
		if lvalue[i] == lastPrime {
			return true, i
		}
	}
	return false, 0
}
