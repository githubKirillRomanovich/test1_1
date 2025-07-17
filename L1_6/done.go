package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Stopped via done channel.")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	close(done)
	time.Sleep(500 * time.Millisecond)
}
