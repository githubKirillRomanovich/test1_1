package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)

	go func() {
		for val := range data {
			fmt.Println("Received:", val)
		}
		fmt.Println("Channel closed, exiting goroutine.")
	}()

	data <- 1
	data <- 2
	close(data)
	time.Sleep(500 * time.Millisecond)
}
