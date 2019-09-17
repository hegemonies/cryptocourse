package test

import (
	DF "cryptocrouse/src/1lab/Diffie-Hellman"
	"cryptocrouse/src/1lab/EuclideanAlgorithm"
	"cryptocrouse/src/1lab/FastExp"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

const (
	MaxNumber = 100000000
	MaxCountTest = 100000
)

func TestFastExpCorrectStable(t *testing.T) {
	var x, y, m int64 = 13, 17, 15

	in := FastExp.FastExp(x, y, m)
	wait := FastExp.NaivePowWithMod(x, y, m)

	if in != wait {
		t.Errorf("Expected %d, got %d", wait, in)
	}
}

func TestFastExpCorrectRand(t *testing.T) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < MaxCountTest; i++ {
		x, y, m := rand.Int63n(MaxNumber) + 1, rand.Int63n(MaxNumber) + 1, rand.Int63n(MaxNumber) + 1

		in := FastExp.FastExp(x, y, m)
		wait := FastExp.NaivePowWithMod(x, y, m)

		if in != wait {
			t.Errorf("i = %d Expected %d, got %d (x = %d, y = %d, m = %d)", i, wait, in, x, y, m)
		}
	}
}

func TestSmallFastExpCorrectRand(t *testing.T) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < MaxCountTest; i++ {
		x, y, m := rand.Uint64() % MaxNumber + 1, rand.Uint64() % MaxNumber + 1, rand.Uint64() % MaxNumber + 1

		in := FastExp.SmallFastExp(x, y, m)
		wait := FastExp.NaivePowWithMod(int64(x), int64(y), int64(m))

		if in != uint64(wait) {
			t.Errorf("i = %d Expected %d, got %d (x = %d, y = %d, m = %d)", i, wait, in, x, y, m)
		}
	}
}

func TestEuclideanAlgoRandStable(t *testing.T) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < MaxCountTest; i++ {
		a, b := rand.Int63n(MaxNumber), rand.Int63n(MaxNumber)

		in := EuclideanAlgorithm.GCD(a, b)
		wait := big.NewInt(0).GCD(big.NewInt(0), big.NewInt(0), big.NewInt(a), big.NewInt(b)).Int64()

		if in != wait {
			t.Errorf("i = %d Expected %d, got %d (a = %d, b = %d)", i, wait, in, a, b)
		}
	}
}

func TestPrimeFunc(t *testing.T) {
	var testingNumber uint64 = 2
	for ; testingNumber < MaxCountTest; testingNumber++ {
		in := DF.IsPrime(testingNumber)
		wait := big.NewInt(int64(testingNumber)).ProbablyPrime(0)

		if in != wait {
			t.Errorf("Expected %v, got %v (testing number = %d)", wait, in, testingNumber)
		}
	}
}

func TestConnectionUser(t *testing.T) {
	cryptoSystem := DF.CryptoSystem{}
	alice := "Alice"
	bob := "Bob"

	for i := 0; i < MaxCountTest; i++ {

		cryptoSystem.Init()
		_ = cryptoSystem.AddUser(alice)
		_ = cryptoSystem.AddUser(bob)
		cryptoSystem.ConnectUsers(alice, bob)

		result := cryptoSystem.CheckConnection(alice, bob)

		if result != true {
			t.Errorf("Expected %v, got %v", true, result)
		}
	}
}