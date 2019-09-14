package Diffie_Hellman

import (
	. "cryptocrouse/src/1lab/FastExp"
	"fmt"
	"math/rand"
	"time"
)

type CryptoSystem struct {
	Users 	map[string]User
	p 		uint64
	g		uint64
}

func (system *CryptoSystem) Init() {
	rand.Seed(time.Now().Unix())
	system.Users = make(map[string]User)
	system.setPrimeNumbers()
}

func (system *CryptoSystem) setPrimeNumbers() {
	q := generatePrimeNumber()
	system.p = 2 * q + 1
	system.g = rand.Uint64() % system.p

	for SmallFastExp(system.g, q, system.p) == 1 {
		if system.g = rand.Uint64() % system.p; system.g < 2 {
			system.g++
		}
	}
}

func (system *CryptoSystem) AddUser(name string) (err error) {
	user := User{}
	user.Name = name
	user.privateKey = system.generatePrivateKey()
	user.PublicKey = system.generatePublicKey(user.privateKey)
	system.Users[name] = user
	return
}

func (system *CryptoSystem) generatePrivateKey() (privateKey uint64) {
	if privateKey = rand.Uint64() % system.p; privateKey < 1 {
		privateKey++
	}
	return
}

func (system *CryptoSystem) generatePublicKey(privateKey uint64) (publicKey uint64) {
	publicKey = SmallFastExp(system.g, privateKey, system.p)
	return
}

func (system *CryptoSystem) PrintUsers() {
	fmt.Println("Name\t\tPrivateKey  \t\tPublicKey")
	for _, user := range system.Users {
		fmt.Printf("%v\t\t%v\t\t%v\n", user.Name, user.privateKey, user.PublicKey)
	}
}