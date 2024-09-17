package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		// Если элемент найден
		if arr[mid] == target {
			return mid
		}

		// Если целевой элемент меньше среднего, ищем в левой части
		if arr[mid] > target {
			right = mid - 1
		} else {
			// Если целевой элемент больше среднего, ищем в правой части
			left = mid + 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7

	result := binarySearch(arr, target)

	if result != -1 {
		fmt.Printf("Элемент %d найден на позиции %d.\n", target, result)
	} else {
		fmt.Printf("Элемент %d не найден.\n", target)
	}
}
