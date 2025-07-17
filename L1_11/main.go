package main

import (
	"fmt"
)

func intersection(a, b []int) []int {
	m := make(map[int]struct{})
	var result []int

	for _, val := range a {
		m[val] = struct{}{}
	}

	for _, val := range b {
		if _, found := m[val]; found {
			result = append(result, val)
			delete(m, val)
		}
	}

	return result
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	inter := intersection(A, B)
	fmt.Println("Пересечение:", inter)
}
