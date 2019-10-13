package ElGamalCode

import (
	"cryptocrouse/src/go/FastExp"
	"cryptocrouse/src/go/FileWrapper"
	"fmt"
	"math/rand"
	"time"
)

type CryptoSystem struct {
	Users map[string]User
	q     uint64
	p     uint64
	g     uint64
}

func (system *CryptoSystem) Init() {
	rand.Seed(time.Now().Unix())
	system.Users = make(map[string]User)
	system.setPrimeNumbers()
}

func (system *CryptoSystem) generateQ() {
	system.q = GeneratePrimeNumber()
}

func (system *CryptoSystem) generateP() {
	system.p = 2 * system.q + 1
	if !IsPrime(system.p) {
		system.generateQ()
		system.generateP()
	}
}

func (system *CryptoSystem) generateG() {
	system.g = rand.Uint64() % system.p
	for FastExp.SmallFastExp(system.g, system.q, system.p) == 1 {
		if system.g = rand.Uint64() % system.p; system.g < 2 {
			system.g++
		}
	}
}

func (system *CryptoSystem) setPrimeNumbers() {
	system.generateQ()
	system.generateP()
	system.generateG()
}

func (system *CryptoSystem) AddUser(name string) (err error) {
	user := User{}
	user.Name = name
	user.GeneratePrivateKey(system.p)
	user.GeneratePublicKey(system.p, system.g)
	system.Users[name] = user
	return
}

func (system *CryptoSystem) PrintUsers() { // refactor this
	fmt.Printf("q = %d\tp = %d\tg = %d\n", system.q, system.p, system.g)
	fmt.Printf("%12s%15s%15s%15s%15s\n", "Name", "Private key", "Public key", "Origin message", "Chunks")
	for _, user := range system.Users {
		fmt.Printf("%12s%15d%15d %v %v\n", user.Name, user.privateKey, user.PublicKey, user.originMessage, user.Chunks)
	}
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

	producerInSystem.originMessage = make([]uint64, len(data))
	producerInSystem.Chunks = make([]Message, 0)

	copy(producerInSystem.originMessage, data)
	producerInSystem.EncodeOriginalMessage(system.g, system.p, consumerInSystem.PublicKey)

	consumerInSystem.originMessage = make([]uint64, len(data))
	consumerInSystem.Chunks = make([]Message, len(data))

	copy(consumerInSystem.Chunks, producerInSystem.Chunks)
	consumerInSystem.DecodeChunks(system.p)

	system.Users[producerName] = producerInSystem
	system.Users[consumerName] = consumerInSystem
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

	data := FileWrapper.GetMessageFromFile(filename)

	producerInSystem.originMessage = make([]uint64, len(data))
	producerInSystem.Chunks = make([]Message, 0)

	copy(producerInSystem.originMessage, data)
	producerInSystem.EncodeOriginalMessage(system.g, system.p, consumerInSystem.PublicKey)

	consumerInSystem.originMessage = make([]uint64, len(data))
	consumerInSystem.Chunks = make([]Message, len(data))

	copy(consumerInSystem.Chunks, producerInSystem.Chunks)
	consumerInSystem.DecodeChunks(system.p)

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

	for i := range producerInSystem.originMessage {
		if producerInSystem.originMessage[i] != consumerInSystem.originMessage[i] {
			return false
		}
	}

	return true
}