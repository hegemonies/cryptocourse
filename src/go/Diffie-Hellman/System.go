package Diffie_Hellman

import (
	. "cryptocrouse/src/go/FastExp"
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

func (system *CryptoSystem) setPrimeNumbers() {
	system.q = GeneratePrimeNumber()
	system.p = 2 * system.q + 1
	for !IsPrime(system.p) {
		system.q = GeneratePrimeNumber()
		system.p = 2 * system.q + 1
	}

	system.g = rand.Uint64() % system.p

	for SmallFastExp(system.g, system.q, system.p) == 1 {
		if system.g = rand.Uint64() % system.p; system.g < 2 {
			system.g++
		}
	}
}

func (system *CryptoSystem) AddUser(name string) (err error) {
	user := User{}
	user.Name = name
	user.GeneratePrivateKey(system.p)
	user.GeneratePublicKey(system.p, system.g)
	system.Users[name] = user
	return
}

func (system *CryptoSystem) PrintUsers() {
	fmt.Printf("q = %d\tp = %d\tg = %d\n", system.q, system.p, system.g)
	fmt.Printf("%12s%15s%15s%15s\n", "Name", "Private key", "Public key", "Secret key")
	for _, user := range system.Users {
		fmt.Printf("%12s%15d%15d%15d\n", user.Name, user.privateKey, user.PublicKey, user.secretKey)
	}
}

func (system *CryptoSystem) ConnectUsers(nameA, nameB string) {
	userA, ok := system.Users[nameA]
	if !ok {
		return // todo: need return error
	}

	userB, ok := system.Users[nameB]
	if !ok {
		return // todo: too
	}

	userA.GenerateSecretKey(userB.PublicKey, system.p)
	userB.GenerateSecretKey(userA.PublicKey, system.p)

	system.Users[nameA] = userA
	system.Users[nameB] = userB
}

func (system *CryptoSystem) CheckConnection(nameA, nameB string) bool {
	userA, ok := system.Users[nameA]
	if !ok {
		return false // todo: need return error
	}

	userB, ok := system.Users[nameB]
	if !ok {
		return false // todo: too
	}

	return userA.secretKey == userB.secretKey
}
