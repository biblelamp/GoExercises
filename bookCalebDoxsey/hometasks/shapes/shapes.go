package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

type Line struct {
	p1, p2 Point
}

type Triangle struct {
	l1, l2, l3 Line
}

type Shapes interface {
	perimeter() float64
}

func (p Point) perimeter() float64 {
	return 0
}

func (l Line) perimeter() float64 {
	return math.Sqrt(math.Pow(l.p1.x - l.p2.x, 2) + math.Pow(l.p1.y - l.p2.y, 2))
}

func (t Triangle) perimeter() float64 {
	return t.l1.perimeter() + t.l2.perimeter() + t.l3.perimeter()
}

func perimeterAll(shapes ...Shapes) {
	for _, s := range shapes {
		fmt.Println(s, "perimetr: ", s.perimeter())
	}
}

func main() {
	p1 := Point{0,0}
	p2 := Point{2, 2}
	p3 := Point{0, 4}
	l1 := Line{p1, p2}
	l2 := Line{p2, p3}
	l3 := Line{p3, p1}
	t := Triangle{l1, l2, l3}

	perimeterAll(p1, l1, l2, l3, t)
}
