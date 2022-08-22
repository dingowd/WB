package main

import (
	"fmt"
	"time"
)

var or = func(channels ...<-chan interface{}) <-chan interface{} {
	done := make(chan interface{})
	for i := 0; i < len(channels); i++ {
		go func(i int) {
			for range channels[i] {
			}
			close(done)
			return
		}(i)
	}
	return done
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(5*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))

}
