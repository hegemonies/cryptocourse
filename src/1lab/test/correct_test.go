package test

import (
	"math/rand"
	"testing"
	"time"

	"../FastExp"
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

	for i := 0; i < 10000000; i++ {
		x, y, m := rand.Int63n(1000000000) + 1, rand.Int63n(1000000000) + 1, rand.Int63n(1000000000) + 1

		in := FastExp.FastExp(x, y, m)
		wait := FastExp.NaivePowWithMod(x, y, m)

		if in != wait {
			t.Errorf("i = %d Expected %d, got %d (x = %d, y = %d, m = %d)", i, wait, in, x, y, m)
		}
	}
}

func TestSmallFastExpCorrectRand(t *testing.T) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < 100000; i++ {
		x, y, m := rand.Uint64() % 1000000000 + 1, rand.Uint64() % 1000000000 + 1, rand.Uint64() % 1000000000 + 1

		in := FastExp.SmallFastExp(x, y, m)
		wait := FastExp.NaivePowWithMod(int64(x), int64(y), int64(m))

		if in != uint64(wait) {
			t.Errorf("i = %d Expected %d, got %d (x = %d, y = %d, m = %d)", i, wait, in, x, y, m)
		}
	}
}