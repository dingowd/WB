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
		go func(v int32) { // запускаем len(arr) горутин
			atomic.AddInt32(&result, v*v) // потокобезопасный подсчет результата с использованием atomic
			wg.Done()
		}(v)
	}
	wg.Wait() // Ждем пока не выполнятся все горутины
	fmt.Fprintln(os.Stdout, result)

	// вариант 2 с mutex
	mu := sync.Mutex{}
	result = 0
	for _, v := range arr { // запускаем len(arr) горутин
		wg.Add(1)
		go func(v int32) {
			mu.Lock() // блокируем все остальные горутины
			result += v * v
			mu.Unlock() // разблокируем все остальные горутины
			wg.Done()
		}(v)
	}
	wg.Wait() // Ждем пока не выполнятся все горутины
	fmt.Fprintln(os.Stdout, result)
}
