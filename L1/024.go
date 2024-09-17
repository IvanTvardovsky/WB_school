package main

import (
	"fmt"
	"math"
)

// Структура Point с инкапсулированными параметрами x и y
type Point struct {
	x float64
	y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) X() float64 {
	return p.x
}

func (p Point) Y() float64 {
	return p.y
}

// Функция для вычисления расстояния между двумя точками
func Distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.X()-p1.X(), 2) + math.Pow(p2.Y()-p1.Y(), 2))
}

func main() {
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)

	fmt.Printf("%.2f\n", Distance(point1, point2))
}
