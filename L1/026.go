package main

import (
	"fmt"
	"strings"
)

func hasUniqueCharacters(s string) bool {
	// Строку к нижнему регистру
	s = strings.ToLower(s)

	// Создаем мапу
	charSet := make(map[rune]struct{})

	for _, char := range s {
		// Проверяем, существует ли символ в множестве
		if _, exists := charSet[char]; exists {
			return false // Если существует => не уникален
		}

		// Добавляем символ в множество
		charSet[char] = struct{}{}
	}

	// Все символы уникальны
	return true
}

func main() {
	testStrings := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
	}

	for _, str := range testStrings {
		fmt.Printf("Строка: \"%s\", Уникальность символов: %v\n", str, hasUniqueCharacters(str))
	}
}
