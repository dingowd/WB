package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

// Sleep используем контекст с таймаутом, при получении значения в канал через ctx.Done просто завершаем работу функции
// всё остальное время не делаем ничего
func Sleep(t time.Duration) {
	ctx, _ := context.WithTimeout(context.Background(), t)
	select {
	case <-ctx.Done():
		return
	}
}

func main() {
	var sec int
	fmt.Fprintln(os.Stdout, "Enter the seconds to sleep")
	fmt.Fscan(os.Stdin, &sec)
	Sleep(time.Duration(sec) * time.Second)
	fmt.Fprintf(os.Stdout, "I slept %v seconds", sec)
}
