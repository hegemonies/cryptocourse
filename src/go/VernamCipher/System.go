package VernamCipher

import (
	"cryptocrouse/src/go/FileWrapper"
	"fmt"
	"math/rand"
	"time"
)

type CryptoSystem struct {
	Users map[string]User
}

func (system *CryptoSystem) Init() {
	rand.Seed(time.Now().Unix())
	system.Users = make(map[string]User)
}

func (system *CryptoSystem) AddUser(name string) (err error) {
	user := User{}
	user.Name = name
	if _, ok := system.Users[name]; !ok {
		system.Users[name] = user
	} else {
		return nil// todo: need add return error
	}
	return nil
}

func (system *CryptoSystem) AddUserWithParams(name string, key []byte) (err error) {
	user := User{}
	user.Name = name
	user.key = key
	if _, ok := system.Users[name]; !ok {
		system.Users[name] = user
	} else {
		return nil// todo: need add return error
	}
	return nil
}

func (system *CryptoSystem) PrintUsers() {
	fmt.Printf("%12s%15s%30s%30sn", "Name", "Origin message", "Encrypt message", "Key")
	for _, user := range system.Users {
		user.PrintUserInfo("%12s%15v%30v%30v\n")
	}
}

func (system *CryptoSystem) SendMessageFromFile(producerName, consumerName, filename string) {
	producerInSystem, ok := system.Users[producerName]
	if !ok {
		return // todo: need return error
	}

	consumerInSystem, ok := system.Users[consumerName]
	if !ok {
		return // todo: too
	}

	producerInSystem.originMessage = FileWrapper.GetMessageFromFileInBytes(filename)
	producerInSystem.GenerateKey()
	producerInSystem.EncryptMessage()

	consumerInSystem.encryptMessage = producerInSystem.encryptMessage
	consumerInSystem.key = producerInSystem.key
	consumerInSystem.DecryptMessage()

	system.Users[producerName] = producerInSystem
	system.Users[consumerName] = consumerInSystem
}

func (system *CryptoSystem) SendMessage(producerName, consumerName string, data []byte) {
	producerInSystem, ok := system.Users[producerName]
	if !ok {
		return // todo: need return error
	}

	consumerInSystem, ok := system.Users[consumerName]
	if !ok {
		return // todo: too
	}

	producerInSystem.originMessage = data
	producerInSystem.GenerateKey()
	producerInSystem.EncryptMessage()

	consumerInSystem.encryptMessage = producerInSystem.encryptMessage
	consumerInSystem.key = producerInSystem.key
	consumerInSystem.DecryptMessage()

	system.Users[producerName] = producerInSystem
	system.Users[consumerName] = consumerInSystem
}

func (system *CryptoSystem) CheckMessage(producerName, consumerName string) (check bool) {
	producerInSystem, ok := system.Users[producerName]
	if !ok {
		return false // todo: need return error
	}

	consumerInSystem, ok := system.Users[consumerName]
	if !ok {
		return false // todo: too
	}

	if len(producerInSystem.originMessage) != len(consumerInSystem.originMessage) {
		return false
	}

	//for i := 0; i < len(producerInSystem.m); i++ {
	for i := range producerInSystem.originMessage {
		if producerInSystem.originMessage[i] != consumerInSystem.originMessage[i] {
			return false
		}
	}

	return true
}
