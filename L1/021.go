package main

import (
	"encoding/json"
	"fmt"
)

// JSONLogger Ожидаемый интерфейс
type JSONLogger interface {
	LogJSON() string
}

// TextLogger Существующая структура, которая выводит текстовые сообщения
type TextLogger struct {
	Message string
}

// LogText Метод TextLogger для вывода текстового сообщения
func (t *TextLogger) LogText() string {
	return t.Message
}

// TextToJSONAdapter Адаптер для TextLogger, который реализует интерфейс JSONLogger
type TextToJSONAdapter struct {
	TextLogger *TextLogger
}

// LogJSON Реализация метода LogJSON, который адаптирует вывод в формате JSON
func (adapter *TextToJSONAdapter) LogJSON() string {
	data, _ := json.Marshal(map[string]string{
		"message": adapter.TextLogger.LogText(),
	})
	return string(data)
}

func main() {
	textLogger := &TextLogger{
		Message: "Hello World!",
	}

	// Используем адаптер для вывода в формате JSON
	adapter := &TextToJSONAdapter{
		TextLogger: textLogger,
	}

	fmt.Println(adapter.LogJSON())
}
