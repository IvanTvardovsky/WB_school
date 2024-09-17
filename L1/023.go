package main

import "fmt"

func removeElement(slice []int, i int) []int {
	// Если индекс вне диапазона
	if i < 0 || i >= len(slice) {
		return slice
	}
	// Удаляем элемент по индексу i, "..." - распаковка
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	i := 2

	fmt.Println(slice)
	slice = removeElement(slice, i)
	fmt.Println(slice)
}
