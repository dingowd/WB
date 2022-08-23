package main

type Burger struct {
	size string
}

func (b *Burger) accept(v Visitor) {
	v.visitForBurger(b)
}

func (b *Burger) getType() string {
	return "burger"
}
