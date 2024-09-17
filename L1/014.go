package main

import (
	"fmt"
)

func determineType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	case bool:
		fmt.Println(v, "bool")
	case chan struct{}:
		fmt.Println(v, "channel")
	default:
		fmt.Println(v, "неизвестный тип")
	}
}

// Можно через reflect

func main() {
	var a interface{}

	a = 42
	determineType(a)

	a = "Hello World!"
	determineType(a)

	a = true
	determineType(a)

	a = make(chan struct{})
	determineType(a)

	a = 3.14
	determineType(a)
}
