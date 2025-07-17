package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d received: %d\n", id, job)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <num_workers>")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Invalid number of workers")
		return
	}

	jobs := make(chan int)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs)
	}

	i := 1
	for {
		jobs <- i
		i++
		time.Sleep(500 * time.Millisecond)
	}
}
