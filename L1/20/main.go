package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Scan() string {
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
	arr := strings.Split(str, " ")
	str = ""
	for i := len(arr) - 1; i >= 0; i-- {
		if i == len(arr)-1 {
			str += arr[i]
		} else {
			str += " " + arr[i]
		}
	}
	fmt.Fprintln(os.Stdout, "Reversed string:")
	fmt.Fprintln(os.Stdout, str)
}
