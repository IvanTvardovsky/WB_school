package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})

	// Добавляем элементы в множество
	for _, word := range words {
		set[word] = struct{}{}
	}

	for word := range set {
		fmt.Println(word)
	}
}
