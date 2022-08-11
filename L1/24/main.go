package main

import (
	"fmt"
	"math"
	"os"
)

type Point struct {
	x int
	y int
}

func New(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

func Distance(p1, p2 Point) float64 {
	dX := float64(p1.x - p2.x)
	dY := float64(p1.y - p2.y)
	dX = math.Pow(dX, 2)
	dY = math.Pow(dY, 2)
	return math.Sqrt(dX + dY)
}
func main() {
	var x, y int
	fmt.Fprint(os.Stdout, "Enter X of point 1:")
	fmt.Fscan(os.Stdin, &x)
	fmt.Fprint(os.Stdout, "Enter Y of point 1:")
	fmt.Fscan(os.Stdin, &y)
	p1 := New(x, y)
	fmt.Fprint(os.Stdout, "Enter X of point 2:")
	fmt.Fscan(os.Stdin, &x)
	fmt.Fprint(os.Stdout, "Enter Y of point 2:")
	fmt.Fscan(os.Stdin, &y)
	p2 := New(x, y)

	fmt.Fprintln(os.Stdout, "Distance between p1 and p2 is:", Distance(p1, p2))
}
