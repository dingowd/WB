package main

type Pizza struct {
	radius int
}

func (p *Pizza) accept(v Visitor) {
	v.visitForPizza(p)
}

func (p *Pizza) getType() string {
	return "pizza"
}
