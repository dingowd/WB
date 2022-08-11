package main

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

type Counter struct {
	c  int32
	wg sync.WaitGroup
}

func (c *Counter) Inc() {
	atomic.AddInt32(&c.c, 1)
	c.wg.Done()
}

func main() {
	var c Counter
	var max int
	fmt.Fprint(os.Stdout, "Enter the max value of counter: ")
	fmt.Fscan(os.Stdin, &max)
	for i := 0; i < max; i++ {
		c.wg.Add(1)
		go c.Inc()
	}
	c.wg.Wait()
	fmt.Fprintln(os.Stdout, c.c)
}
