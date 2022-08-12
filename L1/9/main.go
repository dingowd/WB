package main

import (
	"fmt"
	"os"
	"sync"
)

func Multipler(cIn chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	v := <-cIn
	fmt.Fprintln(os.Stdout, v)
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	cIn := make(chan int)
	wg := new(sync.WaitGroup)
	//cOut := make(chan int)
	for _, v := range arr {
		wg.Add(1)
		cIn <- v
		Multipler(cIn, wg)
	}
	wg.Wait()
	fmt.Fprintln(os.Stdout, "done")
	close(cIn)
}
