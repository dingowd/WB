package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func worker(ch chan int) { // получает произвольное число из канала, в который пишет publisher
	for x := range ch {
		fmt.Fprintln(os.Stdout, "worker received", x)
		// имитация работы с произвольным временем выполнения
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
	var sec int
	fmt.Fprint(os.Stdout, "Enter the time to execute the program in seconds: ")
	fmt.Fscan(os.Stdin, &sec)
	go worker(ch)
	// по истечении времени указанному как второй аргумент context.WithTimeout можно получить сигнал из канала при помощи функции Context.Done()
	stop, cancel := context.WithTimeout(context.Background(), time.Duration(sec)*time.Second)
	defer cancel()
	// сигнал на прерывание пишется в канал, при возможности чтения из этого канала программа завершается
	for {
		select {
		case <-stop.Done():
			close(ch)
			fmt.Fprintln(os.Stdout, "stop the program - time out")
			return
		default:
			publisher(ch)
		}
	}
}
