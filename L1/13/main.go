package main

import (
	"fmt"
	"os"
)

// no comments
func main() {
	a := 1
	b := 2
	fmt.Fprintf(os.Stdout, "Исходные данные:\n %d %d\n", a, b)
	a, b = b, a
	fmt.Fprintf(os.Stdout, "ПОменяли местами:\n %d %d\n", a, b)
}
