package main

import "fmt"

func main() {
	a := 5
	b := 10

	fmt.Printf("a = %d, b = %d\n", a, b)

	// Обмен значениями (под капотом временная переменная не создается)
	// Можно через арифметику
	a, b = b, a

	fmt.Printf("a = %d, b = %d\n", a, b)
}
