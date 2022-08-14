package main

import (
	"fmt"
	"os"
)

func BinarySearch(a []int, e int) *int {
	var mid int
	left := 0
	right := len(a) - 1
	for left <= right {
		mid = (left + right) / 2
		if a[mid] == e {
			return &mid
		}
		if a[mid] > e {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return nil
}

func main() {
	// бинарный поиск возможен только в отсортированном массиве
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
	if result := BinarySearch(arr, 16); result != nil {
		fmt.Fprintln(os.Stdout, *result)
	} else {
		fmt.Fprintln(os.Stdout, result)
	}
}
