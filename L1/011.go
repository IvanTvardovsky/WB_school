package main

import (
	"fmt"
)

// Функция для нахождения пересечения двух множеств
func intersect(setA, setB map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{})

	for key := range setA {
		if _, exists := setB[key]; exists {
			result[key] = struct{}{}
		}
	}

	return result
}

func main() {
	// Определяем два множества
	setA := map[int]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
	}

	setB := map[int]struct{}{
		3: {},
		4: {},
		5: {},
		6: {},
	}

	// Находим пересечение
	intersection := intersect(setA, setB)

	fmt.Println("Пересечение множеств:", intersection)
}
