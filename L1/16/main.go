package main

import (
	"fmt"
	"os"
)

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	split := part(arr)

	QuickSort(arr[:split])
	QuickSort(arr[split:])
}

func part(arr []int) int {
	point := arr[len(arr)/2]

	left := 0
	right := len(arr) - 1

	for {
		for arr[left] < point {
			left++
		}

		for arr[right] > point {
			right--
		}

		if left >= right {
			return right
		}

		arr[left], arr[right] = arr[right], arr[left]
	}
}

func main() {
	arr := []int{2, 5, 8, 1, 4, 9, 3}
	QuickSort(arr)
	fmt.Fprintln(os.Stdout, arr)
}
