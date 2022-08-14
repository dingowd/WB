package main

import (
	"fmt"
	"math/big"
	"os"
)

func Mult(x, y *big.Int) *big.Int {
	res := new(big.Int)
	return res.Mul(x, y)
}

func Div(x, y *big.Int) *big.Int {
	res := new(big.Int)
	return res.Div(x, y)
}

func Add(x, y *big.Int) *big.Int {
	res := new(big.Int)
	return res.Add(x, y)
}

func Sub(x, y *big.Int) *big.Int {
	res := new(big.Int)
	return res.Sub(x, y)
}

// используем пакет big
func main() {
	arg1 := new(big.Int)
	arg2 := new(big.Int)
	var str1, str2 string
	flag := false
	for flag != true { // проверка на то, что ввели именно число
		fmt.Fprint(os.Stdout, "Введите число 1:")
		fmt.Fscan(os.Stdin, &str1)
		_, flag = arg1.SetString(str1, 10)
	}
	flag = false
	for flag != true { // проверка на то, что ввели именно число
		fmt.Fprint(os.Stdout, "Введите число 2:")
		fmt.Fscan(os.Stdin, &str2)
		_, flag = arg2.SetString(str2, 10)
	}
	fmt.Fprintln(os.Stdout, "Умножение: ", Mult(arg1, arg2))
	fmt.Fprintln(os.Stdout, "Деление  : ", Div(arg1, arg2))
	fmt.Fprintln(os.Stdout, "Сложение : ", Add(arg1, arg2))
	fmt.Fprintln(os.Stdout, "Вычитание: ", Sub(arg1, arg2))
}
