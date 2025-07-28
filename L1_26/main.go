package main

import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	seen := make(map[rune]bool)

	for _, char := range strings.ToLower(str) { // приводим к нижнему регистру
		if seen[char] {
			return false // символ уже встречался
		}
		seen[char] = true
	}

	return true
}

func main() {
	tests := []string{
		"abcd",       // true
		"abCdefAaf",  // false (повтор a)
		"aabcd",      // false
		"Hello, Go!", // false (о, l)
		"Uniq123",    // true
	}

	for _, s := range tests {
		fmt.Printf("%q -> %v\n", s, isUnique(s))
	}
}
