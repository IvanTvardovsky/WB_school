package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func worker1(done chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Горутина остановлена через канал")
			return
		default:
			fmt.Println("Работаем...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

/*
func main() {
	done := make(chan bool)
	go worker1(done)

	time.Sleep(2 * time.Second)
	done <- true // Отправляем сигнал остановки
	time.Sleep(1 * time.Second)
}
*/

func worker2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина остановлена через контекст")
			return
		default:
			fmt.Println("Работаем...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

/*
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go worker2(ctx)

	time.Sleep(3 * time.Second) // Ждем, пока контекст завершится
}
*/

func worker3(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("Работаем...")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Горутина завершена через WaitGroup")
}

/*
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go worker3(&wg)

	wg.Wait() // Ожидаем завершения горутины
}

*/

func worker4(stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("Горутина остановлена через таймаут")
			return
		case <-time.After(2 * time.Second):
			fmt.Println("Выполняем задачу")
		}
	}
}

/*
func main() {
	stop := make(chan bool)
	go worker4(stop)

	time.Sleep(3 * time.Second)
	stop <- true // Останавливаем горутину через канал после 3 секунд
	time.Sleep(1 * time.Second)
}

*/

func worker5(stop *int32) {
	for {
		if atomic.LoadInt32(stop) == 1 {
			fmt.Println("Горутина остановлена через флаг")
			return
		}
		fmt.Println("Работаем...")
		time.Sleep(500 * time.Millisecond)
	}
}

/*
func main() {
	var stop int32 = 0
	go worker5(&stop)

	time.Sleep(2 * time.Second)
	atomic.StoreInt32(&stop, 1) // Изменяем флаг
	time.Sleep(1 * time.Second)
}

*/

func worker6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Горутина остановлена через панику:", r)
		}
	}()
	for {
		fmt.Println("Работаем...")
		time.Sleep(500 * time.Millisecond)
	}
}

/*
func main() {
	go worker6()

	time.Sleep(2 * time.Second)
	panic("я паника") // Вызываем панику, которая остановит горутину
}

*/
