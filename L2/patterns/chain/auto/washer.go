package auto

import (
	"fmt"
	"os"
)

type Washer struct {
	Name string
	Next Service
}

func (w *Washer) Serve(c *Car) {
	if !c.ChangeOil {
		fmt.Fprintln(os.Stdout, "Механик еще не сменил масло на", c.Name)
		return
	}
	c.WashCar = true
	fmt.Fprintln(os.Stdout, "Мойщик", w.Name, "помыл машину", c.Name)
}

func (w *Washer) SetNext(s Service) {
	w.Next = s
}
