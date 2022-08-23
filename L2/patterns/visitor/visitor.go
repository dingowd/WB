package main

import (
	"fmt"
	"os"
)

type Visitor interface {
	visitForPizza(p *Pizza)
	vizitForBurger(b *Burger)
	vizitForCoffee(c *Coffee)
}

type VisitorToBuy struct {
	product string
}

func (v *VisitorToBuy) visitForPizza(p *Pizza) {
	fmt.Fprintln(os.Stdout, "I have", p.radius, "cm", v.product)
}

func (v *VisitorToBuy) vizitForBurger(b *Burger) {
	fmt.Fprintln(os.Stdout, "I have", b.size, v.product)
}

func (v *VisitorToBuy) vizitForCoffee(c *Coffee) {
	fmt.Fprintln(os.Stdout, "I have", c.amount, "ml", v.product)
}

type VisitorToCheck struct {
	product string
}

func (v *VisitorToCheck) visitForPizza(p *Pizza) {
	fmt.Fprintln(os.Stdout, "I checked", v.product, "It's really", p.radius, "cm")
}

func (v *VisitorToCheck) vizitForBurger(b *Burger) {
	fmt.Fprintln(os.Stdout, "I checked", v.product, "It is slightly smaller than the", b.size)
}

func (v *VisitorToCheck) vizitForCoffee(c *Coffee) {
	fmt.Fprintln(os.Stdout, "I checked", v.product, " It's really", c.amount, "ml")
}
