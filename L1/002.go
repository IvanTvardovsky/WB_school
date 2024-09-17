package main

import (
	"fmt"
	"sync"
)

func main() {
	sl := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for _, num := range sl {
		wg.Add(1)
		go countSquare(num, &wg)
	}

	wg.Wait()
}

func countSquare(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(num * num)
}
