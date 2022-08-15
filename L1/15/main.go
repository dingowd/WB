package main

import "fmt"

var justString string

// Данная реализация приведет к тому, что в памяти будет храниться переменная v большого объема
// и будет существовать, пока на нее ссылается переменная justString
/*func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}*/

// реализация, которая не ведет к потере памяти
// после выполнения функции someFunc в main память будет очищена от переменной v
func someFunc() {
	v := createHugeString(1 << 10)
	b := make([]rune, 0)
	b = append(b, []rune(v[:100])...)
	justString = string(b)
}

func createHugeString(l int) string {
	a := ""
	for i := 0; i < l; i++ {
		a += "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	}
	return a
}

func main() {
	someFunc()
	fmt.Println(justString)
}
