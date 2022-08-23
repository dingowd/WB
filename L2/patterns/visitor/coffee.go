package main

type Coffee struct {
	amount int
}

func (c *Coffee) accept(v Visitor) {
	v.vizitForCoffee(c)
}

func (p *Coffee) getType() string {
	return "pizza"
}
