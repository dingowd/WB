package main

type Coffee struct {
	amount int
}

func (c *Coffee) accept(v Visitor) {
	v.visitForCoffee(c)
}

func (p *Coffee) getType() string {
	return "coffee"
}
