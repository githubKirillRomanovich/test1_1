package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopped via context.")
				return
			default:
				fmt.Println("Running...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)
}
