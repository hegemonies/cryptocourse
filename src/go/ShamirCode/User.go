package ShamirCode

import (
	Diffie_Hellman "cryptocrouse/src/go/Diffie-Hellman"
	"cryptocrouse/src/go/EuclideanAlgorithm"
	"cryptocrouse/src/go/FastExp"
	"fmt"
	"math/rand"
)

type User struct {
	Name       string
	P          int64
	c          int64
	d          int64
	m          []uint64
}

func (user *User) generateP() {
	q := int64(Diffie_Hellman.GeneratePrimeNumber())
	user.P = 2 * q + 1

	if q < 0 || user.P < 0 {
		user.generateP()
	}

	for !Diffie_Hellman.IsPrime(uint64(user.P)) {
		q = int64(Diffie_Hellman.GeneratePrimeNumber())
		user.P = 2 * q + 1

		if q < 0 || user.P < 0 {
			user.generateP()
		}
	}
}

func (user *User) generateC() {
	var GCD int64 = 0
	for ; GCD != 1; {
		user.c = rand.Int63n(MaxBound)
		if user.c < 2 {
			continue
		}
		GCD, _, _ = EuclideanAlgorithm.GCD(user.c, user.P - 1)
	}
}

func (user *User) generateD() {
	_, _, y := EuclideanAlgorithm.GCD(user.P, user.c)
	if y < 0 {
		user.d = user.P + y
	} else {
		user.d = y
	}
}

func (user *User) GeneratePrivateVariables() {
	user.generateP()
	user.generateC()
	user.generateD()

	for user.d % MaxBound != user.d {
		user.generateP()
		user.generateC()
		user.generateD()
	}
}

func (user *User) GeneratePrivateVariablesWithP(p int64) {
	user.P = p
	user.generateC()
	user.generateD()

	for user.d % MaxBound != user.d {
		user.generateC()
		user.generateD()
	}
}

func (user *User) SetMessage(m []uint64) {
	user.m = m
}

func (user *User) PrintUserInfo(format string) {
	fmt.Printf(format, user.Name, user.P, user.c, user.d, user.m)
}

// encode source user message and return result
func (user *User) ComputeX1() []uint64 {
	return user.computeXPowerC(user.m)
}

func (user *User) ComputeX2(X1 []uint64) []uint64 {
	return user.computeXPowerC(X1)
}

func (user *User) computeXPowerC(X []uint64) []uint64 {
	intermediateX := make([]uint64, 0, len(X))

	for i := 0; i < len(X); i++ {
		intermediateX = append(intermediateX, FastExp.SmallFastExp(X[i], uint64(user.c), uint64(user.P)))
	}

	return intermediateX
}

func (user *User) ComputeX3(X2 []uint64) []uint64 {
	return user.computeXPowerD(X2)
}

func (user *User) ComputeX4(X3 []uint64) []uint64  {
	return user.computeXPowerD(X3)
}

func (user *User) computeXPowerD(X []uint64) []uint64 {
	X_ := make([]uint64, 0, len(X))

	for i := 0; i < len(X); i++ {
		X_ = append(X_, FastExp.SmallFastExp(X[i], uint64(user.d), uint64(user.P)))
	}

	return X_
}

func (user *User) GetMessage() []uint64 {
	return user.m
}

func (user *User) SetC(c int64) {
	user.c = c
}

func (user *User) SetD(d int64) {
	user.d = d
}
