package ShamirCode

import (
	"cryptocrouse/src/go/Diffie-Hellman"
	"cryptocrouse/src/go/FastExp"
	"fmt"
	"math/rand"
)

type User struct {
	Name       string
	p          uint64
	c          uint64
	d          uint64
	m          []uint64
}

func (user *User) GeneratePrivateVariables(p uint64) {
	user.p = Diffie_Hellman.GeneratePrimeNumber()
	user.c = rand.Uint64()
	user.d = rand.Uint64()
	for (user.c * user.d) % (p - 1) != 1 {
		user.c = rand.Uint64()
		user.d = rand.Uint64()
	}
}

func (user *User) SetMessage(m []uint64) {
	user.m = m
}

func (user *User) PrintUserInfo(format string) {
	fmt.Printf(format, user.Name, user.p, user.c, user.d, user.m)
}

func (user *User) SendP() uint64 {
	return user.p
}

func (user *User) SendX1() []uint64 {
	encodeMessages := make([]uint64, 0, len(user.m))

	for i := 0; i < len(user.m); i++ {
		encodeMessages[i] = FastExp.SmallFastExp(user.m[i], user.c, user.p)
	}

	return encodeMessages
}
