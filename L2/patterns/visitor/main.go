package main

import (
	"fmt"
	"os"
)

func main() {
	pizza := &Pizza{
		radius: 25,
	}
	burger := &Burger{
		size: "big",
	}
	coffee := &Coffee{
		amount: 250,
	}

	toBuy := &VisitorToBuy{}
	pizza.accept(toBuy)
	burger.accept(toBuy)
	coffee.accept(toBuy)

	fmt.Fprintln(os.Stdout)

	toCheck := &VisitorToCheck{}
	pizza.accept(toCheck)
	burger.accept(toCheck)
	coffee.accept(toCheck)
}
