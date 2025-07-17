package main

import (
	"fmt"
)

func generate(nums []int, out chan<- int) {
	for _, x := range nums {
		out <- x
	}
	close(out)
}

func multiply(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * 2
	}
	close(out)
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	chanIn := make(chan int)
	chanOut := make(chan int)

	go generate(nums, chanIn)

	go multiply(chanIn, chanOut)

	for result := range chanOut {
		fmt.Println(result)
	}
}
