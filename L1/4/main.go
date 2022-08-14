package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(i int, ch chan int) { // получает произвольное число из канала, в который пишет publisher
	for x := range ch {
		fmt.Fprintln(os.Stdout, "worker", i, "received", x)
		// имитация работы с произвольным временем выполнения от 0 до 999 миллисекунд
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(1000)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}

func publisher(ch chan int) { // пишет в канал произвольное число от 0 до 99
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
	// реализация graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	// сигнал на прерывание пишется в канал, при возможности чтения из этого канала программа завершается
	for {
		select {
		case <-stop:
			close(ch)
			fmt.Fprintln(os.Stdout, "stop the program - interrupted")
			return
		default:
			publisher(ch)
		}
	}
}
