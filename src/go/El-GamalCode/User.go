package El_GamalCode

import (
	"cryptocrouse/src/go/FastExp"
	"math/rand"
)

type User struct {
	Name          string
	PublicKey     uint64 // Y
	privateKey    uint64 // X
	originMessage []uint64
	Chunks        []Message // <message, signature>
}

func (user *User) GeneratePrivateKey(p uint64) {
	if user.privateKey = rand.Uint64() % p; user.privateKey < 1 {
		user.privateKey++
	}
	return
}

func (user *User) GeneratePublicKey(p, g uint64) {
	user.PublicKey = FastExp.SmallFastExp(g, user.privateKey, p)
	return
}

func (user *User) generateRandomK(limit uint64) (k uint64) {
	k = rand.Uint64() % (limit - 1)
	if k < 1 { user.generateRandomK(limit) }
	return
}

func (user *User) EncodeOriginalMessage(g, p uint64) {
	if len(user.originMessage) == 0 {
		return // todo: need return error
	}

	for i := 0; i < len(user.originMessage); i++ {
		k := user.generateRandomK(p)

		chunk := Message{}
		chunk.s = FastExp.SmallFastExp(g, k, p)
		chunk.m = (user.originMessage[i] * FastExp.SmallFastExp(user.PublicKey, k, p)) % p

		user.Chunks = append(user.Chunks, chunk)
	}
}

func (user *User) DecodeChunks(p uint64) {
	user.originMessage = user.originMessage[:0] // clear message
	for i := 0; i < len(user.Chunks); i++ {
		a := user.Chunks[i].s
		b := user.Chunks[i].m
		originMsg := (b * FastExp.SmallFastExp(a, p - 1 - user.privateKey, p)) % p
		user.originMessage = append(user.originMessage, originMsg)
	}
}

func (user *User) GetMessage() []uint64 {
	return user.originMessage
}
