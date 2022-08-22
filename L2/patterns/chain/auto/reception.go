package auto

import (
	"fmt"
	"os"
)

type Reception struct {
	Name string
	Next Service
}

func (r *Reception) Serve(c *Car) {
	if c.GetCar {
		fmt.Fprintln(os.Stdout, "Приёмщик", r.Name, "уже принял машину", c.Name)
		r.Next.Serve(c)
		return
	}
	c.GetCar = true
	fmt.Fprintln(os.Stdout, "Приёмщик", r.Name, "принял машину", c.Name)
	r.Next.Serve(c)
}

func (r *Reception) SetNext(s Service) {
	r.Next = s
}
