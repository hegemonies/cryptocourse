package ShamirCode

import (
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
	fmt.Printf("q = %d\tp = %d\tg = %d\n", system.q, system.p, system.g)
	fmt.Printf("%12s%15s%15s%15s%15s\n", "Name", "P", "C", "D", "Message")
	for _, user := range system.Users {
		user.PrintUserInfo("%12s%15d%15d%15d%15v\n")
	}
}

//func (system *CryptoSystem) ConnectUsers(nameA, nameB string) {
//	userA, ok := system.Users[nameA]
//	if !ok {
//		return // todo: need return error
//	}
//
//	userB, ok := system.Users[nameB]
//	if !ok {
//		return // todo: too
//	}
//
//	userA.GenerateSecretKey(userB.PublicKey, system.p)
//	userB.GenerateSecretKey(userA.PublicKey, system.p)
//
//	system.Users[nameA] = userA
//	system.Users[nameB] = userB
//}
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
