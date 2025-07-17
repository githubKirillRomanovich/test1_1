package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
}

func (c *Counter) Inc() {
	atomic.AddInt64(&c.value, 1)
}

func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	numGoroutines := 1000
	incrementsPerGoroutine := 1000

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Inc()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Итоговое значение счётчика: %d\n", counter.Value())
}
