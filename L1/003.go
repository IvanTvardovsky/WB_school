package main

import (
	"fmt"
	"sync"
)

func main() {
	sl := []int{2, 4, 6, 8, 10}
	ans := 0

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for _, num := range sl {
		wg.Add(1)
		go concurrentSum(num, &ans, &wg, &mu)
	}

	wg.Wait()
	fmt.Println(ans)
}

func concurrentSum(num int, ans *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	*ans += num * num
	mu.Unlock()
}
