package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func worker(ch chan int) {
	for range ch {
		fmt.Fprintln(os.Stdout, "worker received", <-ch)
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
	var sec int
	fmt.Fprint(os.Stdout, "Enter the time to execute program in seconds: ")
	fmt.Fscan(os.Stdin, &sec)
	go worker(ch)
	stop, cancel := context.WithTimeout(context.Background(), time.Duration(sec)*time.Second)
	defer cancel()
	for {
		select {
		case <-stop.Done():
			close(ch)
			fmt.Fprintln(os.Stdout, "stop the program")
			return
		default:
			publisher(ch)
		}
	}
}
