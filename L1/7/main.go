package main

import (
	"fmt"
	"os"
	"sync"
)

type MyMap struct {
	mu sync.Mutex
	m  map[string]int
}

func (m *MyMap) WriteToMap(key string, val int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = val
}

func main() {
	// via mutex
	m := &MyMap{
		m: make(map[string]int),
	}
	keys := []string{"1", "2", "1", "3", "5", "7", "7", "8", "9", "10"}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range keys {
		go m.WriteToMap(v, values[i])
	}
	fmt.Fprintln(os.Stdout, m.m)

	// via sync.Map
	var sm sync.Map
	var wg sync.WaitGroup
	for i, v := range keys {
		wg.Add(1)
		go func(i int, v string) {
			sm.Store(v, values[i])
			wg.Done()
		}(i, v)
	}
	wg.Wait()
	sm.Range(func(k, v interface{}) bool {
		fmt.Println("key:", k, ", val:", v)
		return true
	})

}
