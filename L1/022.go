package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Библиотека для работы с большими числами, но так как 2^20 = 1048576, то большинство операций пройдет и с обычным int64
	// Ещё есть опция с длинной арифметикой
	a := big.NewFloat(1500000)
	b := big.NewFloat(1300000)

	sum := new(big.Float).Add(a, b)
	sub := new(big.Float).Sub(a, b)
	div := new(big.Float).Quo(a, b)
	mul := new(big.Float).Mul(a, b)

	fmt.Println("Sum:", sum)
	fmt.Println("Sub:", sub)
	fmt.Println("Div:", div)
	fmt.Println("Mul:", mul)
}
