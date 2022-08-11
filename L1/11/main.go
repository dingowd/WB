package main

import (
	"fmt"
	"os"
)

func main() {
	set1 := []int{1, 3, 2, 5, 4, 10, 7, 6, 9, 8}
	set2 := []int{7, 5, 12, 28, 3}
	intersection := make([]int, 0)
	// вариант 1 - простой перебор
	for _, s1 := range set1 {
		for _, s2 := range set2 {
			if s1 == s2 {
				intersection = append(intersection, s1)
			}
		}
	}
	fmt.Fprintln(os.Stdout, "вариант 1 - простой перебор")
	fmt.Fprintln(os.Stdout, intersection)

	// вариант 2 - использование map
	intersection = make([]int, 0)
	m1 := make(map[int]struct{})
	m2 := make(map[int]struct{})
	for _, v := range set1 {
		m1[v] = struct{}{}
	}
	for _, v := range set2 {
		m2[v] = struct{}{}
	}
	for k, _ := range m2 {
		_, ok := m1[k]
		if ok {
			intersection = append(intersection, k)
		}
	}
	fmt.Fprintln(os.Stdout, "вариант 2 - использование map")
	fmt.Fprintln(os.Stdout, intersection)
}
