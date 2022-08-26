package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/pflag"
	"io"
	"os"
	"strconv"
	"strings"
)

func revert(a []int) {
	l := len(a)
	for i := 0; i < l/2; i++ {
		a[i], a[l-1-i] = a[l-1-i], a[i]
	}
}

// To process flag p. Can be 1-5,2 or 2,1-5 or 5-1,2,3 as example.
// Interval can be in forward and reverse order
func fProcessing(s string) ([]int, error) {
	res := make([]int, 0)
	spl := strings.Split(s, ",")
	for _, v := range spl {
		if strings.Contains(v, "-") {
			arr := make([]int, 0)
			sep := strings.Split(v, "-")
			if len(sep) != 2 {
				return nil, errors.New("arg fields is invalid")
			}
			var err error
			var left, right int
			var r bool
			left, err = strconv.Atoi(sep[0])
			right, err = strconv.Atoi(sep[1])
			if err != nil {
				return nil, errors.New("arg fields is invalid")
			}
			if left == right {
				arr = append(arr, left)
			} else {
				if left > right {
					r = true
					left, right = right, left
				}
				for i := left; i <= right; i++ {
					arr = append(arr, i)
				}
				if r {
					revert(arr)
				}
				res = append(res, arr...)
				continue
			}
		}
		elem, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("arg fields is invalid")
		}
		res = append(res, elem)
	}
	return res, nil
}

func main() {
	// getting flags
	fields := pflag.StringP("fields", "f", "1", "Fields to output")
	delimiter := pflag.StringP("delimiter", "d", "\t", "Delimiter")
	separated := pflag.BoolP("separated", "s", false, "Separated")
	input := pflag.StringP("input", "i", "", "File to input, default os.Stdin")
	outDelimiter := pflag.StringP("outdelim", "o", " ", "Output delimiter")
	pflag.Parse()

	// set input
	var file io.Reader
	var err error
	if len(*input) > 0 {
		file, err = os.Open(*input)
		if err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
			return
		}
	} else {
		file = os.Stdin
	}
	// getting input
	strArr := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}
		if *separated && !strings.Contains(text, *delimiter) {
			continue
		}
		strArr = append(strArr, text)
	}
	// processing flag p
	f, err := fProcessing(*fields)
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		return
	}
	// filling output
	output := make([]string, 0)
	for _, v := range strArr {
		spl := strings.Split(v, *delimiter)
		collector := make([]string, 0)
		for _, v2 := range f {
			if len(spl) < v2 {
				continue
			}
			collector = append(collector, spl[v2-1])
		}
		str := strings.Join(collector, *outDelimiter)
		output = append(output, str)
	}

	// Print result
	for _, v := range output {
		fmt.Fprintln(os.Stdout, v)
	}
}
