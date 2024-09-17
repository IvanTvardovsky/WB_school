package main

import (
	"fmt"
)

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		// Базовый случай: массив длиной 0 или 1 уже отсортирован
		return arr
	}

	// Выбор опорного элемента
	pivot := arr[0]

	// Подмассивы для элементов меньше и больше опорного элемента
	less := []int{}
	greater := []int{}

	for _, num := range arr[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	// Рекурсивная сортировка подмассивов и объединение результатов
	return append(append(quicksort(less), pivot), quicksort(greater)...)
}

func main() {
	arr := []int{33, 10, 55, 71, 29, 1, 100, 23, 19, 17}

	sortedArr := quicksort(arr)

	fmt.Println("Отсортированный массив:", sortedArr)
}
