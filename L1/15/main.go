package main

import "fmt"

var justString string

// Данная реализация приведет к расходованию большого объема памяти при создании переменной v
/*func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}*/

func someFunc() {
	justString = Huge(1 << 10)[:100]
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
