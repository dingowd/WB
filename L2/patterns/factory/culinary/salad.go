package culinary

import (
	"fmt"
	"os"
)

type Salad struct {
	Type     string
	Products string
	ToSell   string
}

func (s *Salad) BuyProducts() {
	fmt.Fprintln(os.Stdout, "Купить ", s.Products)
}

func (s *Salad) Prepare() {
	fmt.Fprintln(os.Stdout, "Все порезать, посолить, добавить майонез, перемешать")
}

func (s *Salad) Sell() {
	fmt.Fprintln(os.Stdout, s.ToSell)
}

func NewSalad() Culinary {
	s := &Salad{
		Type:     SaladCulinary,
		Products: "огурцы, помидоры, редиска, листья салата, болгарский перец, майонез",
		ToSell:   "Охладить\n",
	}
	fmt.Fprintln(os.Stdout, s.Type)
	return s
}
