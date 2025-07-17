package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	i := 1
	for {
		ch <- i
		i++
		time.Sleep(300 * time.Millisecond)
	}
}

func consumer(ch chan int) {
	for val := range ch {
		fmt.Println("Received:", val)
	}
}

func main() {
	ch := make(chan int)
	timeout := 5 * time.Second

	go producer(ch)
	go consumer(ch)

	<-time.After(timeout)
	fmt.Println("Timeout reached. Stopping...")

	close(ch)

	time.Sleep(1 * time.Second)
}
