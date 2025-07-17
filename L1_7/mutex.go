package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SafeMap) Get(key string) (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val, ok := s.m[key]
	return val, ok
}

func main() {
	safeMap := SafeMap{
		m: make(map[string]int),
	}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			safeMap.Set(key, i)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		if val, ok := safeMap.Get(key); ok {
			fmt.Printf("%s: %d\n", key, val)
		}
	}
}
