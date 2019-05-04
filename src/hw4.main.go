package main

import (
	"fmt"
	"math"
)

// Shape interface for shapes
type Shape interface {
	area() float64
	perimeter() float64
}

// Circle for calculations
type Circle struct {
	radius float64
}

// Square for calculation
type Square struct {
	side float64
}

func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2.)
}

func (circle Circle) perimeter() float64 {
	return math.Pi * circle.radius * 2
}

func (square Square) area() float64 {
	return math.Pow(square.side, 2.)
}

func (square Square) perimeter() float64 {
	return square.side * 4.
}

func sumAll(shapes ...Shape) (areas, perimeters float64) {
	areas, perimeters = 0., 0.
	for _, shape := range shapes {
		areas += shape.area()
		perimeters += shape.perimeter()
	}
	return
}

func main() {
	circle := Circle{10}
	square := Square{10}
	fmt.Println(sumAll(circle, square))
}
