package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	for _, v := range arr {
		wg.Add(1)
		go func(v int) {
			fmt.Fprintln(os.Stdout, v*v)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
