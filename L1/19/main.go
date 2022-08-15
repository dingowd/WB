package main

import (
	"fmt"
	"os"
)

func main() {
	var in string
	fmt.Fprintln(os.Stdout, "Enter the string")
	fmt.Fscan(os.Stdin, &in)
	// Method 1 идем с начала и конца строки к середине
	r1 := []rune(in)
	l := len(r1)
	for i := 0; i < l/2; i++ {
		r1[i], r1[l-i-1] = r1[l-i-1], r1[i]
	}
	fmt.Fprintln(os.Stdout, "Method 1")
	fmt.Fprintln(os.Stdout, string(r1))

	// Method 2 идем с конца строки к началу и добавляем элементы в обратном порядке
	r2 := []rune(in)
	r3 := make([]rune, 0)
	for i := l - 1; i >= 0; i-- {
		r3 = append(r3, r2[i])
	}
	fmt.Fprintln(os.Stdout, "Method 2")
	fmt.Fprintln(os.Stdout, string(r3))
}
