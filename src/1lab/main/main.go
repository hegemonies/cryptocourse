package main

import (
	"../FastExp"
	"fmt"
)

func main() {
	fmt.Println(FastExp.NaivePowWithMod(13, 7, 15))
	fmt.Println(FastExp.FastExp(2, 10, 15))
}
