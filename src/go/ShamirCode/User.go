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
	P          uint64
	c          uint64
	d          uint64
	m          []uint64
}

func (user *User) generateP() {
	q := Diffie_Hellman.GeneratePrimeNumber()
	user.P = 2 * q + 1
	for !Diffie_Hellman.IsPrime(user.P) {
		q = Diffie_Hellman.GeneratePrimeNumber()
		user.P = 2 * q + 1
	}
}

func (user *User) generateC() {
	var GCD uint64 = 0
	for ; GCD != 1; {
		user.c = rand.Uint64()
		if user.c < 2 {
			continue
		}
		GCD, _, _ = EuclideanAlgorithm.GCD(user.c, user.P - 1)
	}
}

func (user *User) generateD() {
	_, _, y := EuclideanAlgorithm.GCD(user.c, user.P)
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

	fmt.Printf("C*D mod p-1 == 1 : %b\n", ((user.c * user.d) % (user.P - 1)) == 1)
}

func (user *User) GeneratePrivateVariablesWithP(p uint64) {
	user.P = p
	user.generateC()
	user.generateD()

	fmt.Printf("C*D mod p-1 == 1 : %b\n", ((user.c * user.d) % (user.P - 1)) == 1)
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
		intermediateX = append(intermediateX, FastExp.SmallFastExp(X[i], user.c, user.P))
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
		X_ = append(X_, FastExp.SmallFastExp(X[i], user.d, user.P))
	}

	return X_
}
