package test

import (
	DF "cryptocrouse/src/go/Diffie-Hellman"
	"cryptocrouse/src/go/EuclideanAlgorithm"
	"cryptocrouse/src/go/FastExp"
	"cryptocrouse/src/go/ShanksAlgorithm"
	"math/big"
	"math/rand"
	"strconv"
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
		EuclideanAlgorithm.GCD(rand.Uint64(), rand.Uint64())
	}
}

func BenchmarkEuclideanAlgoRand(b *testing.B) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < b.N; i++ {
		EuclideanAlgorithm.GCD(rand.Uint64(), rand.Uint64())
	}
}

func BenchmarkDiffieHallmanAlgo(b *testing.B) {
	rand.Seed(time.Now().Unix())

	cryptoSystem := DF.CryptoSystem{}
	cryptoSystem.Init()

	countUsers := b.N

	usersA := make([]string, 0, countUsers)
	for i := 0; i < countUsers; i++ {
		usersA = append(usersA, strconv.Itoa(rand.Int()))
	}

	usersB := make([]string, 0, countUsers)
	for i := 0; i < countUsers; i++ {
		usersB = append(usersB, strconv.Itoa(rand.Int()))
	}

	for i := 0; i < b.N; i++ {
		_ = cryptoSystem.AddUser(usersA[i])
		_ = cryptoSystem.AddUser(usersB[i])
		cryptoSystem.ConnectUsers(usersA[i], usersB[i])
	}
}

func BenchmarkBabyStepGiantStepRand(b *testing.B) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < b.N; i++ {
		a, inX, p := rand.Uint64() %MaxNumber+ 2, rand.Uint64() %MaxNumber+ 1, rand.Uint64() %MaxNumber+ 1
		y := FastExp.SmallFastExp(a, inX, p)

		ShanksAlgorithm.BabyStepGiantStep(a, p, y)
	}
}

func BenchmarkShanksAlgoRand(b *testing.B) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < b.N; i++ {
		a, inX, p := rand.Uint64() %MaxNumber+ 2, rand.Uint64() %MaxNumber+ 1, rand.Uint64() %MaxNumber+ 1
		y := FastExp.SmallFastExp(a, inX, p)

		ShanksAlgorithm.ShanksAlgo(a, p, y)
	}
}
