package main

import (
	"cryptocrouse/src/1lab/ShanksAlgorithm"
	"fmt"
)

func main() {
	//fmt.Println("fast exp naive =", FastExp.NaivePowWithMod(159, 894, 513))
	//fmt.Println("fast exp my    =", FastExp.FastExp(159, 894, 513))
	//
	//fmt.Println("Euclidean my       =", EuclideanAlgorithm.GCD(28, 19))
	//fmt.Println("Euclidean from lib =", big.NewInt(0).GCD(big.NewInt(0),
	//															  big.NewInt(0),
	//															  big.NewInt(28),
	//															  big.NewInt(19)))
	//cryptoSystem := CryptoSystem{}
	//cryptoSystem.Init()
	//cryptoSystem.AddUser("Alice")
	//cryptoSystem.AddUser("Bob")
	//cryptoSystem.ConnectUsers("Alice", "Bob")
	//cryptoSystem.PrintUsers()

	var a, p, y uint64 = 47, 23, 16
	in := ShanksAlgorithm.BabyStepGiantStep(a, p, y)
	fmt.Println(in)
}
