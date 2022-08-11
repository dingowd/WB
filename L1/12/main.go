package main

import (
	"fmt"
	"os"
)

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	groups := make(map[string]struct{})
	for _, v := range str {
		groups[v] = struct{}{}
	}
	out := make([]string, 0)
	for k, _ := range groups {
		out = append(out, k)
	}
	fmt.Fprintln(os.Stdout, out)
}
