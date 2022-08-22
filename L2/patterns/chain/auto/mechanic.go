package auto

import (
	"fmt"
	"os"
)

type Mechanic struct {
	Name string
	Next Service
}

func (m *Mechanic) Serve(c *Car) {
	if !c.GetCar {
		fmt.Fprintln(os.Stdout, "Приемщик не получил машину", c.Name)
		return
	}
	c.ChangeOil = true
	fmt.Fprintln(os.Stdout, "Механик", m.Name, "сменил масло на", c.Name)
	m.Next.Serve(c)
}

func (r *Mechanic) SetNext(s Service) {
	r.Next = s
}
