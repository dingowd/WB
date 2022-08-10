package main

import (
	"fmt"
	"os"
	"time"
)

type Human struct {
	Name string
	Age  uint
}

func (h *Human) PrintName() {
	fmt.Fprintln(os.Stdout, "Name", h.Name)
}

func (h *Human) PrintAge() {
	fmt.Fprintln(os.Stdout, "Age", h.Age)
}

type Action struct {
	Human
	Today time.Time
}

func (a *Action) PrintTime() {
	fmt.Fprintln(os.Stdout, "Today", a.Today.Format("02.01.2006"))
}

func main() {
	a := Action{
		Human: Human{
			Name: "Alex",
			Age:  47,
		},
		Today: time.Now(),
	}
	a.PrintTime()
	a.PrintName()
	a.PrintAge()
}
