package RSA

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
	fmt.Printf("%12s%15s%20s%20s%20s%20s%20s\n", "Name", "C", "D", "P", "Q", "N", "Phi")
	for _, user := range system.Users {
		user.PrintUserInfo("%12s%15d%20d%20d%20d%20d%20d %v\n")
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

	consumerInSystem.GeneratePrivateVariables()

	producerInSystem.SetOrigMessage(FileWrapper.GetMessageFromFile(filename))
	producerInSystem.EncryptMessage(consumerInSystem.D, consumerInSystem.N)
	FileWrapper.WriteToFile("encrypt-" + filename, producerInSystem.GetEncryptMessage())

	consumerInSystem.m = producerInSystem.m
	consumerInSystem.DecryptMessage()
	FileWrapper.WriteToFile("decrypt-" + filename, consumerInSystem.GetOrigMessage())

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

	consumerInSystem.GeneratePrivateVariables()
	producerInSystem.SetOrigMessage(data)
	producerInSystem.EncryptMessage(consumerInSystem.D, consumerInSystem.N)
	consumerInSystem.m = producerInSystem.m
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

	if len(producerInSystem.origM) != len(consumerInSystem.origM) {
		return false
	}

	for i := range producerInSystem.m {
		if producerInSystem.origM[i] != consumerInSystem.origM[i] {
			return false
		}
	}

	return true
}
