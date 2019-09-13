package main

import (
	"../FastExp"
	"fmt"
)


func main() {
	fmt.Println(FastExp.NaivePowWithMod(159, 894, 513))
	fmt.Println(FastExp.FastExp(159, 894, 513))
}
