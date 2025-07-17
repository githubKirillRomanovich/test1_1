package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("Cleanup before Goexit.")
		fmt.Println("Running...")
		runtime.Goexit()
		fmt.Println("Unreachable code")
	}()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Main done.")
}
