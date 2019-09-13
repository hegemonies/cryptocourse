package EuclideanAlgorithm

func EuclideanAlgorithm() {}

func GCD(a, b uint64) uint64 {
	if a <= 0 || b <= 0 {
		return 0
	}

	var q uint64 = 0
	U := make(map[string]uint64)
	V := make(map[string]uint64)
	T := make(map[string]uint64)

	U["GCD"] = a
	U["x"] = 1
	U["y"] = 0

	V["GCD"] = b
	V["x"] = 0
	V["y"] = 1

	for V["GCD"] != 0 {
		q        = U["GCD"] / V["GCD"]
		T["GCD"] = U["GCD"] % V["GCD"]
		T["x"]   = U["x"] - q * V["x"]
		T["y"]   = U["y"] - q * V["y"]
		U = V
		V = T
	}

	return U["GCD"]
}