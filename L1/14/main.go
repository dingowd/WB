package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	a := make([]interface{}, 0)
	a = append(a, 123, "123", true, make(chan int))
	for _, v := range a {
		// simple
		fmt.Fprintf(os.Stdout, "Simple: Type of a: %T\n", v)
		// via reflect
		fmt.Fprintln(os.Stdout, "Via reflect: Type of a:", reflect.TypeOf(v))
		fmt.Fprintln(os.Stdout)
	}
}
