package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d exiting...\n", id)
			return
		case job := <-jobs:
			fmt.Printf("Worker %d received: %d\n", id, job)
		}
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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	jobs := make(chan int)

	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, i, jobs)
	}

	i := 1
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case jobs <- i:
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}

	fmt.Println("Shutting down...")
	time.Sleep(time.Second)
}
