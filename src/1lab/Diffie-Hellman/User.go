package Diffie_Hellman

import (
	. "cryptocrouse/src/1lab/FastExp"
	"math/rand"
)

type User struct {
	Name       string
	PublicKey  uint64
	privateKey uint64
	secretKey  uint64
}

func (user *User) GeneratePrivateKey(p uint64) {
	if user.privateKey = rand.Uint64() % p; user.privateKey < 1 {
		user.privateKey++
	}
	return
}

func (user *User) GeneratePublicKey(p, g uint64) {
	user.PublicKey = SmallFastExp(g, user.privateKey, p)
	return
}

func (user *User) GenerateSecretKey(subscriberPublicKey, p uint64) {
	user.secretKey = SmallFastExp(subscriberPublicKey, user.privateKey, p)
}