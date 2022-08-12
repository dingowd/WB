package main

import "fmt"

var justString string

func someFunc() {
	v := Huge(1 << 10)
	justString = v[:100]
}

func Huge(l int) string {
	a := ""
	for i := 0; i < l; i++ {
		a += "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	}
	return a
}

func main() {
	someFunc()
	fmt.Println(justString)
}
