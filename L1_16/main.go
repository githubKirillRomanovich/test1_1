package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	var left, right []int

	for i, val := range arr {
		if i == pivotIndex {
			continue
		}
		if val <= pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

func main() {
	arr := []int{33, 2, 52, 106, 73, 10, 8, 29}
	sorted := quickSort(arr)
	fmt.Println(sorted)
}
