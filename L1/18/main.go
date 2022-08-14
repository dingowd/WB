package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync/atomic"
	"time"
)

type Counter struct {
	c int32
}

func (c *Counter) Inc() { // потокобезопасное увеличение счетчика на 1 через atomic
	atomic.AddInt32(&c.c, 1)
}

func worker(i int, ch chan struct{}, c *Counter) {
	for range ch {
		// имитация работы с произвольным временем выполнения
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(1000)
		time.Sleep(time.Duration(n) * time.Millisecond)
		c.Inc()
		fmt.Fprintln(os.Stdout, "worker", i, "done job in", n, "milliseconds and set counter to", c.c)
	}
}

func main() {
	ch := make(chan struct{})
	var num int
	c := new(Counter)
	fmt.Fprint(os.Stdout, "Enter the number of workers: ")
	fmt.Fscan(os.Stdin, &num)
	// n воркеров конкурентно инкрементируют счетчик
	for i := 0; i < num; i++ {
		go worker(i+1, ch, c)
	}
	// остановка программы по истечению 5 секунд
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			close(ch)
			fmt.Fprintln(os.Stdout, "Counter now -", c.c)
			fmt.Fprintln(os.Stdout, "stop the program")
			return
		default:
			ch <- struct{}{}
		}
	}
}
