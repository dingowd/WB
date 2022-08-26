package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func match(fixed bool, s string, f string) bool {
	if fixed {
		if strings.Contains(s, f) {
			return true
		} else {
			return false
		}
	} else {
		re, err := regexp.Compile(f)
		if err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
			return false
		}
		str := re.FindString(s)
		if str != "" {
			return true
		} else {
			return false
		}
	}
}

func main() {
	// getting flags
	after := pflag.IntP("after", "A", 0, "печатать +N строк после совпадения")
	before := pflag.IntP("before", "B", 0, "печатать +N строк до совпадения")
	context := pflag.IntP("context", "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	count := pflag.BoolP("count", "c", false, "количество строк")
	ignoreСase := pflag.BoolP("ignore-case", "i", false, "игнорировать регистр")
	invert := pflag.BoolP("invert", "v", false, "вместо совпадения исключать")
	fixed := pflag.BoolP("fixed", "F", false, "точное совпадение со строкой, не паттерн")
	lineNum := pflag.BoolP("linenum", "n", false, "напечатать номер строки")
	pflag.Parse()
	// string to find
	strToFind := pflag.Arg(pflag.NArg() - 2)
	if *ignoreСase {
		strToFind = strings.ToLower(strToFind)
	}
	//open file and read it
	file, err := os.Open(pflag.Arg(pflag.NArg() - 1))
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		return
	}
	// fill slice to process
	strArr := make([]string, 0)
	scanner := bufio.NewScanner(file)
	num := 1
	counter := 0
	for scanner.Scan() {
		text := scanner.Text()
		if *lineNum {
			text = strconv.Itoa(num) + "\t" + text
			num++
		}
		if *count {
			if match(*fixed, text, strToFind) {
				counter++
			}
		}
		if *ignoreСase {
			text = strings.ToLower(text)
		}
		strArr = append(strArr, text)
	}
	// for flag -c printing only number of matched strings
	if *count {
		fmt.Fprintln(os.Stdout, counter)
		return
	}
	// set borders
	var left, right int
	if *context > 0 {
		left, right = *context, *context
	} else {
		left = *before
		right = *after
	}
	// processing array
	out := make([]string, 0)
	// if invert print only unmatched strings
	if *invert {
		for _, v := range strArr {
			if !match(*fixed, v, strToFind) {
				out = append(out, v)
			}
		}
		for _, v := range out {
			fmt.Fprintln(os.Stdout, v)
		}
		return
	}

	for i, v := range strArr {
		if match(*fixed, v, strToFind) {
			// check index
			var leftB, rightB int
			if i-left < 0 {
				leftB = 0
			} else {
				leftB = i - left
			}
			if i+right > len(strArr)-1 {
				rightB = len(strArr) - 1
			} else {
				rightB = i + right
			}
			for i := leftB; i <= rightB; i++ {
				out = append(out, strArr[i])
			}
		}
	}
	for _, v := range out {
		fmt.Fprintln(os.Stdout, v)
	}
}
