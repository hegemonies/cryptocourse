package ShanksAlgorithm

import (
	"cryptocrouse/src/1lab/FastExp"
	"fmt"
	"math"
)

//
//import (
//	"cryptocrouse/src/1lab/FastExp"
//	"fmt"
//	"math"
//)
//
//// (a ** x) mod p = y === lValue = rValue
//// a ** (i * m) = y * (a ** j)
//func BabyStepGiantStep(a, p, y uint64) (x uint64) {
//	m := uint64(math.Pow(float64(p), 0.5)) + 1
//	k := m
//	fmt.Printf("m = %d\tk = %d\n", m, k)
//
//	lvalue := initLvalue(a, p, m, k)
//	rvalue := initRvalue(a, p, y, m)
//
//	var j uint64 = 1
//	var i uint64 = 1
//	// search a ** (i * m) == (a ** j) * y
//	numberIsSearched, fakeI := searchNumber(lvalue, rvalue[0])
//	searched := numberIsSearched
//	for ; searched == false && j < m; lvalue = append(lvalue, nextLvalueElement(a, p, m, i))  {
//		rvalue = append(rvalue, nextRvalueElement(j, a, p, rvalue))
//		for t := 0; t < len(rvalue); t++ {
//			numberIsSearched, fakeI = searchNumber(lvalue, rvalue[t])
//			if numberIsSearched == true {
//				i = fakeI
//				searched = true
//				break
//			}
//		}
//		j++
//		i++
//	}
//
//	j--
//
//	fmt.Printf("i = %d\tj = %d\n", i, j)
//	fmt.Printf("lvalue = %v\n", lvalue)
//	fmt.Printf("rvalue = %v\n", rvalue)
//
//	x = i * m - j
//	return
//}
//
//func initLvalue(a, p, k, m uint64) []uint64 {
//	lvalue := make([]uint64, 0, k)
//	lvalue = append(lvalue, FastExp.SmallFastExp(a, m, p))
//	return lvalue
//}
//
//func initRvalue(a, p, y, m uint64) []uint64 {
//	rvalue := make([]uint64, 0, m)
//	rvalue = append(rvalue, y % p)
//	return rvalue
//}
//
//func nextRvalueElement(j, a, p uint64, rvalue []uint64) uint64 {
//	return (rvalue[j - 1] * a) % p
//}
//
//func nextLvalueElement(a, p, m, i uint64) uint64 {
//	return FastExp.SmallFastExp(a, i * m, p)
//}
//
//func searchNumber(arr []uint64, searchNumber uint64) (searched bool, i uint64) {
//	for i = 0; i < uint64(len(arr)); i++ {
//		if arr[i] == searchNumber {
//			searched = true
//			return
//		}
//	}
//	return false, 0
//}


// Algorithm with using map
// (a ** x) mod p = y === lSet = rSet
// a ** (i * m) = y * (a ** j)
func ShanksAlgo(a, p, y uint64) (x uint64) {
	if p <= y {
		return 0
	}

	m := uint64(math.Pow(float64(p), 0.5) + 0.5)
	k := m
	fmt.Printf("m = %d\tk = %d\n", m, k)

	lSet := make(map[uint64]uint64) // [value]index
	lSet[FastExp.SmallFastExp(a, m, p)] = 1

	rSet := make(map[uint64]uint64) // [index]value
	rSet[0] = y % p

	if _, ok := lSet[rSet[0]]; ok {
		x = 1 * m - 0
		return
	}

	var i uint64 = 2
	var j uint64 = 1

	for i <= k {
		rSet[j] = nextRightElement(rSet[j - 1], a, p)
		lSet[nextLeftElement(a, p, m, i)] = i

		var t uint64 = 0
		for ; t <= j; t++ {
			if currentJ, ok := lSet[rSet[t]]; ok {
				i = t
				j = currentJ
				fmt.Printf("i = %d\tj = %d\n", i, j)
				fmt.Printf("rSet = %v\n", rSet)
				fmt.Printf("lSet = %v\n", lSet)
				x = i * m - j
				return
			}
		}

		i++
		j++
	}

	fmt.Printf("i = %d\tj = %d\n", i, j)
	//fmt.Printf("rSet = %v\n", rSet)
	//fmt.Printf("lSet = %v\n", lSet)

	return
}

func nextRightElement(rvalue, a, p uint64) uint64 {
	return (rvalue * a) % p
}

func nextLeftElement(a, p, m, i uint64) uint64 {
	return FastExp.SmallFastExp(a, i * m, p)
}


func ShanksAlgo2(a, p, y uint64) (x uint64) {
	if p <= y {
		return 0
	}

	m := uint64(math.Pow(float64(p), 0.5) + 0.5)
	k := m
	//fmt.Printf("m = %d\tk = %d\n", m, k)

	lSet := make(map[uint64]uint64) // [value]index
	//rSet := make(map[uint64]uint64) // [index]value
	//rSet[0] = y % p
	//rValue := y % p

	var i uint64 = 1

	for ; i <= k; i++ {
		lSet[FastExp.SmallFastExp(a, i * m, p)] = i
	}

	var j uint64 = 0

	for ; j < m; j++ {
		//rSet[j] = (rSet[j - 1] * a) % p
		//rValue = (rValue * a) % p
		rValue := (FastExp.SmallFastExp(a, j, p) * y) % p

		if currentJ, ok := lSet[rValue]; ok {
			j = currentJ
			x = i * m - j
			return
		}
	}

	return
}

func ShanksAlgo3(a, p, y uint64) (x uint64) {
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
