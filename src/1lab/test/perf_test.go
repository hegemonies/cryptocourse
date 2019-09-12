package test

import (
	"math/big"
	"math/rand"
	"time"

	"testing"
	"../FastExp"
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
	x := big.NewInt(13)
	y := big.NewInt(17)
	m := big.NewInt(15)

	for i := 0; i < b.N; i++ {
		big.NewInt(0).Exp(x, y, m)
	}
}

func BenchmarkFastExpRand(b *testing.B) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < b.N; i++ {
		x, y, m := rand.Int63n(50) + 1, rand.Int63n(50)+ 1, rand.Int63n(50) + 1
		FastExp.FastExp(x, y, m)
	}
}
