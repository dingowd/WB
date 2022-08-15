package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

// С помощью sync.WaitGroup
func goroutine1(wg *sync.WaitGroup) {
	fmt.Fprintln(os.Stdout, "goroutine1 done")
	wg.Done()
}

// С помощью канала
func goroutine2(ch chan struct{}) {
	for {
		select {
		case <-ch:
			fmt.Fprintln(os.Stdout, "goroutine2 done")
			return
		default:
		}
	}
}
func goroutine3(ch chan struct{}) {
	defer func() {
		fmt.Fprintln(os.Stdout, "goroutine3 done")
	}()
	for range ch {
	}
}

// С помощью context.Context
func goroutine4(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Fprintln(os.Stdout, "goroutine4 done")
			return
		}
	}
}

// с помощью panic
func goroutine5(wg *sync.WaitGroup) {
	defer wg.Done()
	n := 0
	for {
		time.Sleep(1 * time.Second)
		if n == 10 {
			panic("I am in panic")
		}
		n++
	}
}

func main() {
	// С помощью sync.WaitGroup
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go goroutine1(wg)
	wg.Wait()

	// С помощью канала
	ch1 := make(chan struct{})
	go goroutine2(ch1)

	ch2 := make(chan struct{})
	go goroutine3(ch2)
	time.Sleep(3 * time.Second)
	ch1 <- struct{}{}
	close(ch1)
	close(ch2)

	// С помощью context.Context
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	go goroutine4(ctx)

	// с помощью panic
	wg.Add(1)
	go goroutine5(wg)
	wg.Wait()
	time.Sleep(4 * time.Second)
	fmt.Fprintln(os.Stdout, "main done") // никогда не выполнится потому что горутина 5 обвалит приложение в панику
}
