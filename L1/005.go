package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	N := 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(N)*time.Second)
	defer cancel()

	// Создаем канал для передачи данных
	dataChan := make(chan int)

	// Горутина для записи данных в канал
	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done():
				// Завершаем отправку данных при завершении контекста
				fmt.Println("Завершение отправки данных")
				close(dataChan)
				return
			case dataChan <- i:
				fmt.Println("Отправлено в канал:", i)
				i++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Горутина для чтения данных из канала
	go func() {
		for num := range dataChan {
			fmt.Printf("Прочитано из канала: %d\n", num)
		}
	}()

	// Ждем завершения контекста
	<-ctx.Done()
	fmt.Println("Программа завершена")
}
