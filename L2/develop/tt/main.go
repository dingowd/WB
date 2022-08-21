package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

type Keys struct {
	K int
	N bool
	R bool
	U bool
}

func main() {
	k := pflag.Int64P("kolumn", "k", 0, "num of column")
	n := pflag.BoolP("numeric", "n", false, "numeric sort")
	r := pflag.BoolP("r", "r", false, "numeric sort")
	u := pflag.BoolP("u", "u", false, "numeric sort")

	pflag.Parse()
	fmt.Println(*k, *n, *r, *u)

}
