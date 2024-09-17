package main

import (
	"context"
	"fmt"
	"time"
)

func SleepWithContext(ctx context.Context, duration time.Duration) error {
	// Создаем новый контекст с таймаутом
	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel() // Отменяем контекст по завершении функции

	// Ждем завершения контекста
	select {
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	fmt.Println("Начало")

	// Создаем базовый контекст
	ctx := context.Background()

	// Ожидаем 2 секунды
	if err := SleepWithContext(ctx, 2*time.Second); err != nil {
		if err.Error() != "context deadline exceeded" {
			fmt.Println(err)
		} else {
			fmt.Println("всё ОК")
		}
	}

	fmt.Println("Конец")
}
