package main

import (
	"fmt"
	"os"
)

func toChan1(nums []int, out chan int) {
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
}

func fromChan1toChan2(in <-chan int, out chan int) {
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
}

// классический пайплайн
// т.к. по заданию даны 2 канала, то инициализируем их в main,
// но закрываем их в функциях, потому что именно они являются писателями в переданные им каналы
func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch1 := make(chan int)
	ch2 := make(chan int)

	toChan1(arr, ch1)
	fromChan1toChan2(ch1, ch2)

	for range arr {
		fmt.Fprintln(os.Stdout, <-ch2)
	}
}
