package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseRunes(runes []rune, start, end int) {
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}

func reverseWordsInPlace(s string) string {
	runes := []rune(s)
	n := len(runes)

	reverseRunes(runes, 0, n-1)

	start := 0
	for i := 0; i <= n; i++ {
		if i == n || runes[i] == ' ' {
			reverseRunes(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

func main() {
	fmt.Print("Введите строку: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		result := reverseWordsInPlace(input)
		fmt.Println("Результат:", result)
	}
}
