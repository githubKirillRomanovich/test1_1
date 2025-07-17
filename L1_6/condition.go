package main

import (
	"fmt"
	"time"
)

func main() {
	stop := false

	go func() {
		for !stop {
			fmt.Println("Running...")
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Println("Stopped by flag.")
	}()

	time.Sleep(1 * time.Second)
	stop = true
	time.Sleep(500 * time.Millisecond)
}
