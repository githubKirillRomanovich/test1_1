package main

import (
	"fmt"
)

func uniqueStrings(items []string) []string {
	set := make(map[string]struct{})
	for _, item := range items {
		set[item] = struct{}{}
	}

	result := make([]string, 0, len(set))
	for key := range set {
		result = append(result, key)
	}
	return result
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	unique := uniqueStrings(words)
	fmt.Println("Уникальные слова:", unique)
}
