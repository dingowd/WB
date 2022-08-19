package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func isAnagramm(first, second string) bool {
	one := strings.Split(first, "")
	two := strings.Split(second, "")
	for i := 0; i < len(one); i++ {
		for j := 0; j < len(two); j++ {
			if two[j] == "" {
				continue
			}
			if one[i] == two[j] {
				two[j] = ""
			}
		}
	}
	for _, v := range two {
		if v != "" {
			return false
		}
	}
	return true
}

func main() {
	arr := []string{"тяпка", "пятка", "Пятак", "листок", "столик", "слиток", "чтотоеще", "столик"}
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.ToLower(arr[i])
	}
	m := make(map[string]map[string]struct{}) // создаем map of map, чтобы сразу исключить повторяющиеся элементы
	for i := 0; i < len(arr); i++ {
		if arr[i] == "" {
			continue
		}
		m[arr[i]] = make(map[string]struct{})
		m[arr[i]][arr[i]] = struct{}{}
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == "" {
				continue
			}
			if isAnagramm(arr[i], arr[j]) {
				m[arr[i]][arr[j]] = struct{}{}
				arr[j] = ""
			}
		}
		arr[i] = ""
	}
	out := make(map[string][]string) // заполняем исходящее множество, проверяем на количество элементов
	for key, _ := range m {
		for k, _ := range m[key] {
			if len(m[key]) > 1 {
				out[key] = append(out[key], k)
			}
		}
	}
	for key, _ := range out {
		sort.Strings(out[key]) // сортируем исходящее множество
	}
	fmt.Fprintln(os.Stdout, out)
}
