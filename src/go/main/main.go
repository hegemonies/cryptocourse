package main

import (
	Diffie_Hellman "cryptocrouse/src/go/Diffie-Hellman"
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

	//var a, p, y uint64 = 47, 23, 16
	//
	//in1 := ShanksAlgorithm.BabyStepGiantStep(a, p, y)
	//fmt.Println(in1)
	//
	//in2 := ShanksAlgorithm.ShanksAlgo(a, p, y)
	//fmt.Println(in2)
	//
	//in3 := ShanksAlgorithm.ShanksAlgo(2, 3, 2)
	//fmt.Println(in3)


	//filename := "test_data.png"
	//var p uint64 = 23
	//FileWrapper.WriteToFile("1" + filename, FileWrapper.GetMessageFromFileByP(filename, p), p)

	//var p uint64 = 1007849279
	//fmt.Println(p)
	//fmt.Println(int64(p) - 1)
	//fmt.Println((107042119 * 859590983) % 1661422823)
	//fmt.Println((107042119 * 859590983) % 1661422823 - 1)
	//fmt.Println((107042119 * 859590983) % (1661422823 - 1))

	var c, d, p uint64

	c, d, p = 759828549, 963602659, 1954085507

	fmt.Println(Diffie_Hellman.IsPrime(1791745079))

	fmt.Println((c * d) % (p - 1))

}
