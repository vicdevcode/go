package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

type Triangle struct {
	p1, p2, p3 Point
}

func (t *Triangle) length(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func (t *Triangle) area() float64 {
	a := t.length(t.p1, t.p2)
	b := t.length(t.p2, t.p3)
	c := t.length(t.p3, t.p1)
	p := (a + b + c) / 2
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func main() {
	t := Triangle{
		p1: Point{0, 0},
		p2: Point{3, 0},
		p3: Point{3, 4},
	}

	fmt.Printf("Площадь треугольника: %.2f\n", t.area())
}
