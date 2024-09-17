package main

import (
	"fmt"
)

func reverseString(input string) string {
	// Преобразуем строку в слайс рун для корректной обработки unicode
	runes := []rune(input)

	// Инвертируем порядок рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	input := "главрыба"
	result := reverseString(input)
	fmt.Println(result)
}
