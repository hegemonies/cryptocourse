package RSA

import (
	diffiehellman "cryptocrouse/src/go/Diffie-Hellman"
	"cryptocrouse/src/go/EuclideanAlgorithm"
	"cryptocrouse/src/go/FastExp"
	"fmt"
	"math/rand"
)

type User struct {
	Name       string
	c          uint64
	D          uint64
	p          uint64
	q          uint64
	N          uint64
	phi        uint64
	origM      []uint64
	m          []uint64
}

func (user *User) generateP() {
	for {
		q := diffiehellman.GeneratePrimeNumberWithBounds(MaxBound)
		user.p = 2 * q + 1

		if diffiehellman.IsPrime(user.p) {
			if user.p >= MinBound && user.p <= MaxBound {
				break
			}
		}
	}
}

func (user *User) generateQ() {
	user.q = diffiehellman.GeneratePrimeNumberWithBounds(MaxBound)
}

func (user *User) generatePhi() {
	user.phi = (user.p - 1) * (user.q - 1)
}

func (user *User) generateN() {
	user.N = user.p * user.q
}

func (user *User) generateC() {
	_, _, y := EuclideanAlgorithm.GCD(int64(user.phi), int64(user.D))
	if y < 0 {
		user.c = uint64(int64(user.p) + y)
	} else {
		user.c = uint64(y)
	}
}

func (user *User) generateD() {
	for {
		user.D = rand.Uint64() % MaxBound

		if user.D < 2 || user.D >= user.phi {
			continue
		}

		GCD, _, _ := EuclideanAlgorithm.GCD(int64(user.phi), int64(user.D))
		if GCD == 1 {
			break
		}
	}
}

func (user *User) GeneratePrivateVariables() {
	for {
		user.generateP()
		user.generateQ()
		user.generateN()
		user.generatePhi()
		user.generateD()
		user.generateC()

		if user.D <= MaxBound && ((user.c * user.D) % (user.phi)) == 1 {
			break
		}
	}
}

func (user *User) PrintUserInfo(format string) {
	fmt.Printf(format, user.Name, user.c, user.D, user.p, user.q, user.N, user.phi, user.origM)
}

func (user *User) EncryptMessage(d, n uint64) {
	user.m = make([]uint64, len(user.origM))
	for i := 0; i < len(user.m); i++ {
		user.m[i] = FastExp.SmallFastExp(user.origM[i], d, n)
	}
}

func (user *User) DecryptMessage() {
	user.origM = make([]uint64, len(user.m))
	for i := 0; i < len(user.m); i++ {
		user.origM[i] = FastExp.SmallFastExp(user.m[i], user.c, user.N)
	}
}

func (user *User) GetOrigMessage() []uint64 {
	return user.origM
}

func (user *User) SetOrigMessage(m []uint64) {
	user.origM = m
}

func (user *User) GetEncryptMessage() []uint64 {
	return user.m
}

func (user *User) SetEncryptMessage(m []uint64) {
	user.m = m
}
