package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const numOfWorkers = 3

func worker(id int, ctx context.Context, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d заканчивает работу\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				return // Канал закрыт
			}
			fmt.Printf("Воркер %d начал джобу: %d\n", id, job)
			time.Sleep(time.Second)
			fmt.Printf("Воркер %d сделал джобу: %d\n", id, job)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	sl := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30}

	jobs := make(chan int, 2)

	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go worker(i, ctx, jobs, &wg)
	}

	// Горутина для обработки сигналов
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		cancel() // Отмена контекста
	}()

	// Отправка данных в канал
	for _, num := range sl {
		select {
		case <-ctx.Done():
			// Если контекст завершен, выходим из цикла и не отправляем больше данные
			fmt.Println("Прекращение отправки данных из-за завершения контекста")
			break
		case jobs <- num:
		}
	}

	// Закрытие канала после завершения отправки данных
	close(jobs)
	wg.Wait()
	fmt.Println("Работа завершена")
}
