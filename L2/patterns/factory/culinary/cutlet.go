package culinary

import (
	"fmt"
	"os"
)

type Cutlet struct {
	Type     string
	Products string
	ToSell   string
}

func (c *Cutlet) BuyProducts() {
	fmt.Fprintln(os.Stdout, "Купить ", c.Products)
}

func (c *Cutlet) Prepare() {
	fmt.Fprintln(os.Stdout, "Прокрутить фарш, посолить, поперчить, перемешать, слепить котлеты")
}

func (c *Cutlet) Sell() {
	fmt.Fprintln(os.Stdout, c.ToSell)
}

func NewCutlet() Culinary {
	c := &Cutlet{
		Type:     CutletCulinary,
		Products: "говядина, свинина, соль, перец, панировка",
		ToSell:   "Заморозить\n",
	}
	fmt.Fprintln(os.Stdout, c.Type)
	return c
}
