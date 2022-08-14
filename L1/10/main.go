package main

import (
	"fmt"
	"os"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)
	for _, v := range temps {
		key := int(v) / 10 * 10              //(-25 / 10 * 10) = 20 = (-27 /10 *10)
		groups[key] = append(groups[key], v) // добавляем в слайс map с ключом key элемент исходного слайса
	}
	fmt.Fprintln(os.Stdout, groups)
}
