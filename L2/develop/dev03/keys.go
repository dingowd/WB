package main

import (
	"strconv"
	"strings"
	"unicode"
)

type Keys struct {
	K int
	N bool
	R bool
	U bool
}

func SetKeys(arg string) Keys {
	arg = strings.ToLower(arg)
	args := []rune(arg)
	k := Keys{
		K: 0,
		N: false,
		R: false,
		U: false,
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
				k.K = 0
				break
			}
			if !unicode.IsDigit(args[i+1]) {
				k.K = 0
			} else {
				k.K, _ = strconv.Atoi(string(args[i+1]))
				k.K--
				i++
			}
		case 110: // n
			k.N = true
		case 114: // r
			k.R = true
		case 117: // u
			k.U = true
		}
		i++
	}
	return k
}
