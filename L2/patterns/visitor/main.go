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

	toBuy := &VisitorToBuy{
		product: "pizza",
	}
	pizza.accept(toBuy)
	toBuy = &VisitorToBuy{
		product: "burger",
	}
	burger.accept(toBuy)
	toBuy = &VisitorToBuy{
		product: "coffee",
	}
	coffee.accept(toBuy)

	fmt.Fprintln(os.Stdout)

	toCheck := &VisitorToCheck{
		product: "pizza",
	}
	pizza.accept(toCheck)
	toCheck = &VisitorToCheck{
		product: "burger",
	}
	burger.accept(toCheck)
	toCheck = &VisitorToCheck{
		product: "coffee",
	}
	coffee.accept(toCheck)
}
