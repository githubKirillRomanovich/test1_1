package main

import (
	"fmt"
	"sync"
)

func squareWorker(n int, wg *sync.WaitGroup, results chan int) {
	defer wg.Done()
	square := n * n
	results <- square
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan int, len(numbers))
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go squareWorker(num, &wg, results)
	}

	wg.Wait()
	close(results)

	for square := range results {
		fmt.Println(square)
	}
}
