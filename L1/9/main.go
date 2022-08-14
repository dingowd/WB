package main

import (
	"fmt"
	"os"
)

func toChan1(nums []int, out chan int) <-chan int {
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func fromChan1toChan2(in <-chan int, out chan int) <-chan int {
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// классический пайплайн
func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch1 := make(chan int)
	ch2 := make(chan int)

	toMult := toChan1(arr, ch1)
	toPrint := fromChan1toChan2(toMult, ch2)

	for range arr {
		fmt.Fprintln(os.Stdout, <-toPrint)
	}
}
