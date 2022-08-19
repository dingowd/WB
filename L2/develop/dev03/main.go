package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*type Keys struct {
	k int
	n bool
	r bool
	u bool
}

func SetKeys(arg string) Keys {
	args := []rune(arg)
	k := Keys{
		k: 0,
		n: false,
		r: false,
		u: false,
	}
	if len(arg) == 0 {
		return k
	}
	i := 0
	l := len(args)
	for i < l {
		switch args[i] {
		case 107: // k
			if (i + 2) > l {
				k.k = 0
				break
			}
			if !unicode.IsDigit(args[i+1]) {
				k.k = 0
			} else {
				k.k, _ = strconv.Atoi(string(args[i+1]))
				k.k--
				i++
			}
		case 110: // n
			k.n = true
		case 114: // r
			k.r = true
		case 117: // u
			k.u = true
		}
		i++
	}
	return k
}*/

type SortElem struct {
	Elem string
	Str  string
}

type ForSort []SortElem

func (f ForSort) Len() int { return len(f) }

func (f ForSort) Less(i, j int) bool { return f[i].Elem < f[j].Elem }

func (f ForSort) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

type ForSortRev []SortElem

func (f ForSortRev) Len() int { return len(f) }

func (f ForSortRev) Less(i, j int) bool { return f[i].Elem > f[j].Elem }

func (f ForSortRev) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

type ForSortNum []SortElem

func (f ForSortNum) Len() int { return len(f) }

// Less если выбранные элементы нельзя преобразовать в число, то произойдет построчная сортировка
func (f ForSortNum) Less(i, j int) bool {
	left, errL := strconv.Atoi(f[i].Elem)
	right, errR := strconv.Atoi(f[j].Elem)
	if errL != nil || errR != nil {
		return f[i].Elem < f[j].Elem
	}
	return left < right
}

func (f ForSortNum) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

type ForSortNumRev []SortElem

func (f ForSortNumRev) Len() int { return len(f) }

// Less если выбранные элементы нельзя преобразовать в число, то произойдет построчная сортировка
func (f ForSortNumRev) Less(i, j int) bool {
	left, errL := strconv.Atoi(f[i].Elem)
	right, errR := strconv.Atoi(f[j].Elem)
	if errL != nil || errR != nil {
		return f[i].Elem > f[j].Elem
	}
	return left > right
}

func (f ForSortNumRev) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

func main() {
	// получаем аргументы командной строки и получаем из нее ключи при помощи функции SetKeys
	keys := SetKeys(strings.ReplaceAll(os.Args[1:][0], "-", ""))
	strArr := make([]string, 0)
	// чтение из файла
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		strArr = append(strArr, fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	defer file.Close()

	num := keys.K // номер столбца, по которому будет происходить сортировка
	delim := " "
	f := make(ForSort, 0)
	for _, v := range strArr { // Заполняем структуру для сортировки
		a := strings.Split(v, delim)
		var s SortElem
		s.Elem = a[num]
		s.Str = v
		f = append(f, s)
	}
	if keys.U { // убираем повторяющиеся значения по ключу u
		m := make(map[string]string)
		for _, v := range f {
			m[v.Str] = v.Elem
		}
		f = make(ForSort, 0)
		for k, v := range m {
			var s SortElem
			s.Elem = v
			s.Str = k
			f = append(f, s)
		}
	}
	if keys.R { // проверяем на обратный порядок по ключу r
		if keys.N {
			sort.Sort(ForSortNumRev(f)) // сортировка по числовому значению в обратном порядке
		} else {
			sort.Sort(ForSortRev(f)) // обычная сортировка в обратном порядке
		}
	} else {
		if keys.N {
			sort.Sort(ForSortNum(f)) // прямая сортировка по числовому значению
		} else {
			sort.Sort(ForSort(f)) // прямая обычная сортировка
		}
	}
	result := make([]string, 0)

	for _, v := range f {
		result = append(result, v.Str)
	}

	fileOut, errF := os.Create("out.txt")
	defer fileOut.Close()
	if errF != nil {
		log.Fatalf("Error when opening file: %s", errF)
	}
	for _, v := range result {

		io.WriteString(fileOut, (v + "\n"))
	}
}
