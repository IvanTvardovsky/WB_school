package main

import (
	"fmt"
	"sync"
)

// Counter Структура счетчика
type Counter struct {
	mu    sync.Mutex
	count int
}

// Increment Метод для инкрементации счетчика
func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// Value Метод для получения текущего значения счетчика
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	var wg sync.WaitGroup

	counter := &Counter{}

	numGoroutines := 100
	wg.Add(numGoroutines)

	// Запуск конкурентных горутин, которые инкрементируют счетчик
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	fmt.Println(counter.Value())
}
