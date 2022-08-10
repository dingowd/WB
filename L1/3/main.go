package main

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

func main() {
	// вариант 1 с atomic
	arr := []int32{2, 4, 6, 8, 10}
	var result int32
	wg := sync.WaitGroup{}
	for _, v := range arr {
		wg.Add(1)
		go func(v int32) {
			atomic.AddInt32(&result, v*v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Fprintln(os.Stdout, result)

	// вариант 2 с mutex
	mu := sync.Mutex{}
	result = 0
	for _, v := range arr {
		wg.Add(1)
		go func(v int32) {
			mu.Lock()
			result += v * v
			mu.Unlock()
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Fprintln(os.Stdout, result)
}
