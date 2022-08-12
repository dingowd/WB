package main

import (
	"fmt"
	"os"
	"sync"
)

func Mult(chIN chan int, wg *sync.WaitGroup) {
	// var in int
	defer wg.Done()
	for range chIN {
		fmt.Fprintln(os.Stdout, <-chIN)
		/*		in = <- chIN
				chOUT <- in * 2*/
	}
}

func Print(chOUT chan int) {
	for range chOUT {
		fmt.Fprintln(os.Stdout, <-chOUT)
	}
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	chIN := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	//chOUT := make(chan int)
	//go Print(chOUT)
	go Mult(chIN, wg)
	for _, v := range arr {
		chIN <- v
	}
	wg.Wait()
	close(chIN)
	//close(chOUT)
}
