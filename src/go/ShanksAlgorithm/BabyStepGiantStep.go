package ShanksAlgorithm

import "cryptocrouse/src/go/FastExp"

func BabyStepGiantStep(a, p, y uint64) (x uint64) {
	if p <= y {
		return 0
	}

	var m, k uint64

	k = p / 2
	m = k + 1
	k++

	var i, j uint64

	rvalueList := make([]uint64, 0, k)

	var prev uint64
	prev = y % p
	rvalueList = append(rvalueList, prev)

	for j = 1; j < m; j++ {
		prev = (prev * a) % p
		rvalueList = append(rvalueList, prev)
	}

	needExit := false
	for i = 1; (i <= k) && needExit != true; i++ {
		prev = FastExp.SmallFastExp(a, i * m, p)

		for j = 0; j < uint64(len(rvalueList)); j++ {
			if prev == rvalueList[j] {
				needExit = true
				break
			}
		}
	}

	i--

	x = i * m - j

	return
}
