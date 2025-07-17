package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			m.Store(key, i)
		}(i)
	}

	wg.Wait()

	m.Range(func(key, value any) bool {
		fmt.Printf("%s: %d\n", key, value)
		return true
	})
}
