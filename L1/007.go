package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Структура с map и мьютексом для синхронизации
type SafeMap struct {
	mu    sync.Mutex
	mymap map[int]string
}

// Метод для безопасной записи в map
func (sm *SafeMap) Set(key int, value string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.mymap[key] = value
}

// Метод для безопасного чтения из map
func (sm *SafeMap) Get(key int) (string, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	value, ok := sm.mymap[key]
	return value, ok
}

func main() {
	sm := SafeMap{
		mymap: make(map[int]string),
	}

	var wg sync.WaitGroup

	// Количество горутин для записи данных
	numWorkers := 5

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Set(i, strconv.Itoa(i*i))
		}(i)
	}

	// Ждем завершения всех горутин
	wg.Wait()

	for i := 0; i < numWorkers; i++ {
		if value, ok := sm.Get(i); ok {
			fmt.Printf("Ключ: %d, Значение: %s\n", i, value)
		} else {
			fmt.Printf("Ключ %d не найден\n", i)
		}
	}
}
