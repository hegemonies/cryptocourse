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

func (system *CryptoSystem) AddUserWithParams(name string, c, d, p uint64) (err error) {
	user := User{}
	user.Name = name
	user.P = p
	user.c = c
	user.d = d
	if _, ok := system.Users[name]; !ok {
		system.Users[name] = user
	} else {
		return nil// todo: need add return error
	}
	return nil
}

func (system *CryptoSystem) PrintUsers() {
	fmt.Printf("%12s%15s%30s%30s %s\n", "Name", "P", "C", "D", "Message")
	for _, user := range system.Users {
		user.PrintUserInfo("%12s%15d%30d%30d %v\n")
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

	producerInSystem.SetMessage(FileWrapper.GetMessageFromFileByP(filename, uint64(producerInSystem.P)))

	X1 := producerInSystem.ComputeX1()
	X2 := consumerInSystem.ComputeX2(X1)
	X3 := producerInSystem.ComputeX3(X2)
	refMessage := consumerInSystem.ComputeX4(X3)
	consumerInSystem.SetMessage(refMessage)

	system.Users[producerName] = producerInSystem
	system.Users[consumerName] = consumerInSystem
}

func (system *CryptoSystem) SendMessage(producerName, consumerName string, data []uint64) {
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

	producerInSystem.SetMessage(data)

	X1 := producerInSystem.ComputeX1()
	X2 := consumerInSystem.ComputeX2(X1)
	X3 := producerInSystem.ComputeX3(X2)
	refMessage := consumerInSystem.ComputeX4(X3)
	consumerInSystem.SetMessage(refMessage)

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

	if len(producerInSystem.m) != len(consumerInSystem.m) {
		return false
	}

	//for i := 0; i < len(producerInSystem.m); i++ {
	for i := range producerInSystem.m {
		if producerInSystem.m[i] != consumerInSystem.m[i] {
			return false
		}
	}

	return true
}

func (system *CryptoSystem) UpdateUser(name string, user User) {
	u, _ := system.Users[name]
	u.P = user.P
	u.c = user.c
	u.d = user.d
}
