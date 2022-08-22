package culinary

import (
	"fmt"
	"os"
)

const (
	BunCulinary    = "булочка"
	SaladCulinary  = "салат"
	CutletCulinary = "котлета"
)

type Culinary interface {
	BuyProducts()
	Prepare()
	Sell()
}

func New(product string) Culinary {
	switch product {
	case BunCulinary:
		return NewBun()
	case SaladCulinary:
		return NewSalad()
	case CutletCulinary:
		return NewCutlet()
	default:
		fmt.Fprintln(os.Stdout, "Что-то новенькое:", product)
		return nil
	}
}
