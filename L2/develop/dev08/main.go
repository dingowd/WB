package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Shell v 1.0")
	cur, _ := os.Getwd()
	fmt.Fprintln(os.Stdout, cur)
	os.Chdir("..")
	cur, _ = os.Getwd()
	fmt.Fprintln(os.Stdout, cur)
}
