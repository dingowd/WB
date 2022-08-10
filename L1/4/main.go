package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(i int, ch chan int) {
	for range ch {
		fmt.Fprintln(os.Stdout, "worker", i, "received", <-ch)
		// имитация работы с произвольным временем выполнения
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(1000)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}

func publisher(ch chan int) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	ch <- n
}

func main() {
	ch := make(chan int)
	var num int
	fmt.Fprint(os.Stdout, "Enter the number of workers: ")
	fmt.Fscan(os.Stdin, &num)
	for i := 0; i < num; i++ {
		go worker(i+1, ch)
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case <-stop:
			close(ch)
			fmt.Fprintln(os.Stdout, "stop the program")
			return
		default:
			publisher(ch)
		}
	}
}
