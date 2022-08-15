package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Scan() string { // функция написана для того, чтобы из считывать строку из Stdin с пробелами
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Enter error:", err)
	}
	return in.Text()
}

func main() {
	var str string
	fmt.Fprintln(os.Stdout, "Enter the string to reverse:")
	str = Scan()
	arr := strings.Split(str, " ") // разделяем слова по пробелу и создаем слайс, где слова являются элементами данного слайса
	l := len(arr)
	// Method 1 идем с начала и конца строки к середине
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
	str = strings.Join(arr, " ")
	fmt.Fprintln(os.Stdout, "Method 1 Reversed string:")
	fmt.Fprintln(os.Stdout, str)
	// Method 2 идем по массиву в обратном порядке
	str = arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		str += " " + arr[i]
	}
	fmt.Fprintln(os.Stdout, "Method 1 Reversed again:")
	fmt.Fprintln(os.Stdout, str)
}
