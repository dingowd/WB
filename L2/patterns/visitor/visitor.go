package main

import (
	"fmt"
	"os"
)

type Visitor interface {
	visitForPizza(p *Pizza)
	visitForBurger(b *Burger)
	visitForCoffee(c *Coffee)
}

type VisitorToBuy struct {
}

func (v *VisitorToBuy) visitForPizza(p *Pizza) {
	fmt.Fprintln(os.Stdout, "I have", p.radius, "cm", p.getType())
}

func (v *VisitorToBuy) visitForBurger(b *Burger) {
	fmt.Fprintln(os.Stdout, "I have", b.size, b.getType())
}

func (v *VisitorToBuy) visitForCoffee(c *Coffee) {
	fmt.Fprintln(os.Stdout, "I have", c.amount, "ml", c.getType())
}

type VisitorToCheck struct {
}

func (v *VisitorToCheck) visitForPizza(p *Pizza) {
	fmt.Fprintln(os.Stdout, "I checked", p.getType(), "It's really", p.radius, "cm")
}

func (v *VisitorToCheck) visitForBurger(b *Burger) {
	fmt.Fprintln(os.Stdout, "I checked", b.getType(), "It is slightly smaller than the", b.size)
}

func (v *VisitorToCheck) visitForCoffee(c *Coffee) {
	fmt.Fprintln(os.Stdout, "I checked", c.getType(), "It's really", c.amount, "ml")
}
