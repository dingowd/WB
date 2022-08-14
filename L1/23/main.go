package main

import (
	"fmt"
	"os"
)

func main() {
	// стандартная операция (медленная, не меняет порядок элементов)
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var idx int
	fmt.Fprintln(os.Stdout, "Какой элемент массива", arr, "удалить?")
	fmt.Fscan(os.Stdin, &idx)
	if idx >= len(arr) || idx < 0 {
		fmt.Fprintln(os.Stdout, "Index out of range")
		return
	}
	arr = append(arr[:idx], arr[idx+1:]...)
	fmt.Fprintln(os.Stdout, "Правильно\n", arr)

	// более быстрая (меняет порядок элементов)
	arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr[idx] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	fmt.Fprintln(os.Stdout, "Быстро\n", arr)
}
