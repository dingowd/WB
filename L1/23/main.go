package main

import (
	"fmt"
	"os"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var idx int
	fmt.Fscan(os.Stdin, &idx)
	if idx >= len(arr) || idx < 0 {
		fmt.Fprintln(os.Stdout, "Index out of range")
		return
	}
	arr = append(arr[:idx], arr[idx+1:]...)
	fmt.Fprintln(os.Stdout, arr)
}
