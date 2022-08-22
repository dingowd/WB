package culinary

import (
	"fmt"
	"os"
)

type Bun struct {
	Type     string
	Products string
	ToSell   string
}

func (b *Bun) BuyProducts() {
	fmt.Fprintln(os.Stdout, "Купить ", b.Products)
}

func (b *Bun) Prepare() {
	fmt.Fprintln(os.Stdout, "Замесить тесто, придать форму, испечь, посыпать сахаром")
}

func (b *Bun) Sell() {
	fmt.Fprintln(os.Stdout, b.ToSell)
}

func NewBun() Culinary {
	b := &Bun{
		Type:     BunCulinary,
		Products: "мука, соль, сахар, яйца, дрожжи, молоко",
		ToSell:   "Продавать только горячей\n",
	}
	fmt.Fprintln(os.Stdout, b.Type)
	return b
}
