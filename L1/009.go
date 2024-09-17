package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	inputChannel := make(chan int)
	outputChannel := make(chan int)

	// Горутина для чтения из массива и записи в inputChannel
	go func() {
		for _, num := range numbers {
			inputChannel <- num
		}
		// Закрываем канал после отправки всех чисел
		close(inputChannel)
	}()

	// Горутина для обработки чисел из inputChannel и записи результата в outputChannel
	go func() {
		for num := range inputChannel {
			outputChannel <- num * 2
		}
		// Закрываем канал после обработки всех чисел
		close(outputChannel)
	}()

	// Чтение из outputChannel и вывод результата в stdout
	for result := range outputChannel {
		fmt.Println(result)
	}
}
