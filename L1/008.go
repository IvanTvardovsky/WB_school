package main

import (
	"fmt"
)

// Установить i-й бит в 1 (побитовый or)
func setBit(n int64, i int) int64 {
	return n | (1 << i)
}

// Установить i-й бит в 0 (побитовый and с побитовой инверсией)
func clearBit(n int64, i int) int64 {
	return n &^ (1 << i)
}

func main() {
	var num int64 = 42
	i := 2 // Это бит считая справа, начиная с 0

	fmt.Printf("Исходное число: %d (в двоичном: %b)\n", num, num)

	num = setBit(num, i)
	fmt.Printf("После установки %d-го бита в 1: %d (в двоичном: %b)\n", i, num, num)

	num = clearBit(num, i)
	fmt.Printf("После установки %d-го бита в 0: %d (в двоичном: %b)\n", i, num, num)
}
