package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	str = "snow dog sun"
	arr := strings.Split(str, " ")
	str = ""
	for i := len(arr) - 1; i >= 0; i-- {
		if i == len(arr)-1 {
			str += arr[i]
		} else {
			str += " " + arr[i]
		}
	}
	fmt.Fprintln(os.Stdout, str)
}
