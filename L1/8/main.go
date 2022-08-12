package main

import (
	"fmt"
	"os"
)

func main() {
	var num, res int64
	bNum := 65
	fmt.Fprint(os.Stdout, "Введите исходное число:")
	fmt.Fscan(os.Stdin, &num)
	for bNum > 63 || bNum < 0 {
		fmt.Fprint(os.Stdout, "Введите номер бита от 0 до 63:")
		fmt.Fscan(os.Stdin, &bNum)
	}
	var bit int64 = 2
	for bit != 0 && bit != 1 {
		fmt.Fprint(os.Stdout, "установить 0 или 1 ?:")
		fmt.Fscan(os.Stdin, &bit)
	}
	fmt.Fprintf(os.Stdout, "%064b, %v - исходное число\n", num, num)
	// Способ 1 - через получение остатка от деления
	if num != -(1 << (64 - 1)) {
		os1 := num % (1 << bNum)
		os2 := num % (1 << (bNum + 1))
		res = num - os2
		if bit == 0 {
			res += os1
		} else {
			if num >= 0 {
				res = res + os1 + 1<<bNum
			} else {
				res = res + os1 - 1<<bNum
			}
		}
	} else {
		res = res >> (64 - bNum - 1)
	}
	fmt.Fprintln(os.Stdout, "Способ 1 - через получение остатка от деления")
	fmt.Fprintf(os.Stdout, "%064b, %v - результат\n", res, res)

	// Способ 2 - с помощью побитовых операций
	res = 0
	if num < 0 {
		res -= num
	} else {
		res = num
	}
	x := (res >> (bNum)) & 1 // Получение бита с номером bNum
	if x != bit {
		res = res ^ (1 << bNum) // Переключение бита с номером bNum
	}
	if num < 0 {
		res = -res
	}
	fmt.Fprintln(os.Stdout, "Способ 2 - с помощью побитовых операций")
	fmt.Fprintf(os.Stdout, "%064b, %v - результат\n", res, res)
}
