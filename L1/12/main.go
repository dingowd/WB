package main

import (
	"fmt"
	"os"
)

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	groups := make(map[string]struct{})
	for _, v := range str { // пишем в map, в ней не может быть повторяющихся ключей
		groups[v] = struct{}{}
	}
	res := make([]string, 0)
	for k, _ := range groups { // добавляем все ключи в массив результата
		res = append(res, k)
	}
	fmt.Fprintln(os.Stdout, res)
}
