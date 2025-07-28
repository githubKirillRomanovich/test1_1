package main

import (
	"fmt"
)

func removeIndex(slice []int, i int) []int {
	// Проверка границ
	if i < 0 || i >= len(slice) {
		return slice
	}

	// Сдвиг хвоста влево на одну позицию
	copy(slice[i:], slice[i+1:])

	// Обнуляем последний элемент (для избежания утечки памяти)
	slice[len(slice)-1] = 0

	// Возвращаем слайс с уменьшенной длиной
	return slice[:len(slice)-1]
}

func main() {
	s := []int{10, 20, 30, 40, 50}
	fmt.Println("Before:", s)

	s = removeIndex(s, 2) // Удаляем элемент с индексом 2 (30)
	fmt.Println("After:", s)
}
