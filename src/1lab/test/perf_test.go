package test

import (
	"cryptocrouse/src/1lab/EuclideanAlgorithm"
	"cryptocrouse/src/1lab/FastExp"
	"math/big"
	"math/rand"
	"time"

	"testing"
)

func BenchmarkFastExpStable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastExp.FastExp(13, 15, 17)
	}
}

func BenchmarkFastExpStable2(b *testing.B) {
	var x, y, m int64 = 13, 15, 17
	for i := 0; i < b.N; i++ {
		FastExp.FastExp(x, y, m)
	}
}

func BenchmarkStdExpStable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := big.NewInt(13)
		y := big.NewInt(17)
		m := big.NewInt(15)

		big.NewInt(0).Exp(x, y, m)
	}
}

func BenchmarkFastExpRand(b *testing.B) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < b.N; i++ {
		x, y, m := rand.Int63n(1000000000) + 1, rand.Int63n(1000000000)+ 1, rand.Int63n(1000000000) + 1
		FastExp.FastExp(x, y, m)
	}
}

func BenchmarkSmallFastExpRand(b *testing.B) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < b.N; i++ {
		x, y, m := rand.Uint64() % 1000000000 + 1, rand.Uint64() % 1000000000 + 1, rand.Uint64() % 1000000000 + 1
		FastExp.SmallFastExp(x, y, m)
	}
}

func BenchmarkFastExpStableRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x, y, m := rand.Int63n(1000000000) + 1, rand.Int63n(1000000000) + 1, rand.Int63n(1000000000) + 1
		FastExp.FastExp(x, y, m)
	}
}

func BenchmarkSmallFastExpStableRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x, y, m := rand.Uint64() % 1000000000 + 1, rand.Uint64() % 1000000000 + 1, rand.Uint64() % 1000000000 + 1
		FastExp.SmallFastExp(x, y, m)
	}
}

func BenchmarkEuclideanAlgoStable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EuclideanAlgorithm.GCD(28, 19)
	}
}

func BenchmarkEuclideanAlgoRandStable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EuclideanAlgorithm.GCD(rand.Int63(), rand.Int63())
	}
}

func BenchmarkEuclideanAlgoRand(b *testing.B) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < b.N; i++ {
		EuclideanAlgorithm.GCD(rand.Int63(), rand.Int63())
	}
}
