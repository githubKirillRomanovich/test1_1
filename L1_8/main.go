package main

import (
	"fmt"
)

func SetBit(value int64, i int, bit int) int64 {
	if bit == 1 {
		return value | (1 << i)
	} else {
		return value &^ (1 << i)
	}
}

func main() {
	var x int64 = 5
	i := 1
	bit := 0

	result := SetBit(x, i, bit)
	fmt.Printf("Исходное число: %d (%04b)\n", x, x)
	fmt.Printf("Результат:      %d (%04b)\n", result, result)
}
