package main

import (
	"../EuclideanAlgorithm"
	"../FastExp"
	"fmt"
	"math/big"
)


func main() {
	fmt.Println("fast exp naive =", FastExp.NaivePowWithMod(159, 894, 513))
	fmt.Println("fast exp my    =", FastExp.FastExp(159, 894, 513))

	fmt.Println("Euclidean my       =", EuclideanAlgorithm.GCD(28, 19))
	fmt.Println("Euclidean from lib =", big.NewInt(0).GCD(big.NewInt(0),
																  big.NewInt(0),
																  big.NewInt(28),
																  big.NewInt(19)))
}
