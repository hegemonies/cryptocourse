package ShanksAlgorithm
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


