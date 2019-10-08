package ShamirCode

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

func (system *CryptoSystem) PrintUsers() {
	fmt.Printf("%12s%15s%15s%15s%15s\n", "Name", "P", "C", "D", "Message")
	for _, user := range system.Users {
		user.PrintUserInfo("%12s%15d%15d%15d%15v\n")
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

	producerInSystem.GeneratePrivateVariables()
	consumerInSystem.GeneratePrivateVariablesWithP(producerInSystem.P)

	producerInSystem.SetMessage(FileWrapper.GetMessageFromFileByP(filename, producerInSystem.P))

	X1 := producerInSystem.ComputeX1()
	X2 := consumerInSystem.ComputeX2(X1)
	X3 := producerInSystem.ComputeX3(X2)
	refMessage := consumerInSystem.ComputeX4(X3)
	consumerInSystem.SetMessage(refMessage)

	system.Users[producerName] = producerInSystem
	system.Users[consumerName] = consumerInSystem
}
//
//func (system *CryptoSystem) CheckConnection(nameA, nameB string) bool {
//	userA, ok := system.Users[nameA]
//	if !ok {
//		return false // todo: need return error
//	}
//
//	userB, ok := system.Users[nameB]
//	if !ok {
//		return false // todo: too
//	}
//
//	return userA.secretKey == userB.secretKey
//}
