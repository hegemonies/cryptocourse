package test

import (
	"math/rand"
	"testing"

	"../FastExp"
)

func TestFastExpCorrectStable(t *testing.T) {
	var x, y, m int64 = 13, 17, 15

	myFastExp := FastExp.FastExp(x, y, m)
	expFromMathLib := FastExp.NaivePowWithMod(x, y, m)

	if myFastExp != expFromMathLib {
		t.Errorf("Expected %d, got %d", expFromMathLib, myFastExp)
	}
}

func TestFastExpCorrectRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		x, y, m := rand.Int63n(50) + 1, rand.Int63n(50)+ 1, rand.Int63n(50) + 1

		myFastExp := FastExp.FastExp(x, y, m)
		expFromMathLib := FastExp.NaivePowWithMod(x, y, m)

		if myFastExp != expFromMathLib {
			t.Errorf("Expected %d, got %d", expFromMathLib, myFastExp)
		}
	}
}