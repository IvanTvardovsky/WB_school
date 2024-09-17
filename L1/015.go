package main

import (
	"fmt"
)

func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFunc() string {
	v := createHugeString(1 << 10)
	return v[:100] // Возвращаем только нужную часть - создается новая строка копируя первые 100 символов
}

var justString string

func main() {
	justString = someFunc()
	fmt.Println(justString)
}

// someFunc возвращает строку вместо того, чтобы присваивать её глобальной переменной
// Это позволяет избежать ненужного хранения больших объектов в памяти
