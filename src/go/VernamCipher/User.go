package VernamCipher

import (
	"crypto/rand"
	"fmt"
)

type User struct {
	Name           string
	originMessage  []byte
	encryptMessage []byte
	key            []byte
}

func (user *User) GenerateKey() {
	if len(user.originMessage) == 0 {
		return // todo: return error
	}

	user.key = make([]byte, len(user.originMessage))
	_, _ = rand.Read(user.key)
}

func (user *User) EncryptMessage() {
	if len(user.originMessage) == 0 || len(user.key) == 0 {
		return // todo: return error
	}

	user.encryptMessage = make([]byte, len(user.originMessage))
	for i := 0; i < len(user.originMessage); i++ {
		user.encryptMessage[i] = user.originMessage[i] ^ user.key[i]
	}
}

func (user *User) DecryptMessage() {
	if len(user.encryptMessage) == 0 || len(user.key) == 0 {
		return // todo: return error
	}

	user.originMessage = make([]byte, len(user.encryptMessage))
	for i := 0; i < len(user.encryptMessage); i++ {
		user.originMessage[i] = user.encryptMessage[i] ^ user.key[i]
	}
}

func (user *User) SetMessage(message []byte) {
	user.originMessage = message
}

func (user *User) PrintUserInfo(format string) {
	fmt.Printf(format, user.Name, user.originMessage, user.encryptMessage, user.key)
}

func (user *User) GetOriginMessage() []byte {
	return user.originMessage
}
