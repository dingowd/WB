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

// пишем в map элемент с ключом key и значением val
func (m *MyMap) WriteToMap(key string, val int, wg *sync.WaitGroup) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = val
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	// via mutex
	m := &MyMap{
		m: make(map[string]int),
	}
	keys := []string{"1", "2", "1", "3", "5", "7", "7", "8", "9", "10"}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range keys { // запускаем len(keys) горутин
		wg.Add(1)
		go m.WriteToMap(v, values[i], &wg)
	}

	// via sync.Map
	var sm sync.Map
	for i, v := range keys { // запускаем len(keys) горутин
		wg.Add(1)
		go func(i int, v string) {
			sm.Store(v, values[i])
			wg.Done()
		}(i, v)
	}
	wg.Wait() // ждём выполнения всех горутин
	fmt.Fprintln(os.Stdout, "via mutex")
	for k, v := range m.m {
		fmt.Fprint(os.Stdout, k, ":", v, " ")
	}
	fmt.Fprintln(os.Stdout, "")
	fmt.Fprintln(os.Stdout, "via sync.Map")
	sm.Range(func(k, v interface{}) bool {
		fmt.Fprint(os.Stdout, k, ":", v, " ")
		return true
	})
}
