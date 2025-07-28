package main

import (
	"fmt"
	"math"
)

// Структура Point с приватными полями
type Point struct {
	x float64
	y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Метод для вычисления расстояния между двумя точками
func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(3.0, 4.0)
	p2 := NewPoint(0.0, 0.0)

	fmt.Printf("Расстояние между точками: %.2f\n", p1.Distance(p2))
}
