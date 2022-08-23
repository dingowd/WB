package main

type Burger struct {
	size string
}

func (b *Burger) accept(v Visitor) {
	v.vizitForBurger(b)
}

func (b *Burger) getType() string {
	return "pizza"
}
