package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func isDgt(r rune) bool {
	return unicode.IsDigit(r)
}

func isStringCorrect(s string) bool {
	wr := []rune(s)
	if s == "" {
		return true
	}
	if isDgt(wr[0]) {
		return false
	}
	ln := len(wr)
	for i := 1; i < ln-1; i++ {
		if isDgt(wr[i]) && isDgt(wr[i+1]) && string(wr[i-1]) != `\` {
			return false
		}
	}
	return true
}

func Unpack(inStr string) (string, error) {
	var outStr strings.Builder
	if !isStringCorrect(inStr) {
		return "", ErrInvalidString
	}
	wr := []rune(inStr)
	ln := len(wr)
	if ln == 0 {
		return "", nil
	}
	for i := 0; i < ln-1; i++ {
		if isDgt(wr[i+1]) {
			n, _ := strconv.Atoi(string(wr[i+1]))
			outStr.WriteString(strings.Repeat(string(wr[i]), n))
		} else if !isDgt(wr[i]) {
			outStr.WriteRune(wr[i])
		}
	}
	outStr.WriteRune(wr[ln-1])
	return outStr.String(), nil
}
