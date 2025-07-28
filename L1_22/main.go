package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Инициализация больших чисел
	a := big.NewInt(0)
	b := big.NewInt(0)

	// Устанавливаем значения > 2^20 (например, 2^40 = 1_099_511_627_776)
	a.SetString("1099511627776", 10) // 2^40
	b.SetString("1048576", 10)       // 2^20

	// Сложение
	sum := big.NewInt(0).Add(a, b)

	// Вычитание
	diff := big.NewInt(0).Sub(a, b)

	// Умножение
	mul := big.NewInt(0).Mul(a, b)

	// Деление
	div := big.NewInt(0).Div(a, b) // целочисленное деление

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("a + b =", sum)
	fmt.Println("a - b =", diff)
	fmt.Println("a * b =", mul)
	fmt.Println("a / b =", div)
}
