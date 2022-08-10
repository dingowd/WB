package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	fmt.Fscan(os.Stdin, &s)
	s = strings.ToLower(s)
	// вариант 1 с использованием map
	m := make(map[string]bool)
	for _, v := range s {
		m[string(v)] = true
	}
	if len(s) == len(m) {
		fmt.Fprintln(os.Stdout, "true")
	} else {
		fmt.Fprintln(os.Stdout, "false")
	}
	// вариант 2 с использованием цикла
	result := true
outer:
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				result = false
				break outer
			}
		}
	}
	fmt.Fprintln(os.Stdout, result)
}
