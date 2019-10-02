package EuclideanAlgorithm

func GCD(a, b int64) (int64, int64, int64) {
	if a <= 0 || b <= 0 {
		return 0, 0, 0
	}

	var q int64 = 0
	U := make(map[string]int64)
	V := make(map[string]int64)
	T := make(map[string]int64)

	U["GCD"] = a
	U["x"] = 1
	U["y"] = 0

	V["GCD"] = b
	V["x"] = 0
	V["y"] = 1

	for V["GCD"] != 0 {
		q = U["GCD"] / V["GCD"]
		T["GCD"] = U["GCD"] % V["GCD"]
		T["x"] = U["x"] - q * V["x"]
		T["y"] = U["y"] - q * V["y"]
		swapMap(U, V)
		swapMap(V, T)
	}

	return U["GCD"], U["x"], U["y"]
}

func swapMap(src, dst map[string]int64) {
	if len(src) != len(dst) {
		return
	}

	tmpMap := make(map[string]int64)

	for key, value := range src {
		tmpMap[key] = value
	}

	for key, value := range dst {
		src[key] = value
	}

	for key, value := range tmpMap {
		dst[key] = value
	}
}
