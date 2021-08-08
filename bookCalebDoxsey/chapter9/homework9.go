package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float32
}

type Square struct {
	x, y, d float32
}

type Shape interface {
	area() float32
}

type Person struct {
	name string
}

type Android struct {
	Person
	model string
}

func main() {
	c := Circle{0,0,5}
	s := Square{0,0, 3}
	fmt.Println(c)
	fmt.Println(circleAreaDefault(0,0, 5))
	fmt.Println(circleArea(c))
	fmt.Println(c.area())
	fmt.Println(s.area())
	fmt.Println(areas(c, s))

	a := new(Android)
	a.Person.name = "Android"
	a.Person.Talk()

	a.Talk()
}

func circleAreaDefault(x, y, r float32) float32 {
	return math.Pi * r * r
}

func circleArea(c Circle) float32 {
	return math.Pi * c.r * c.r
}

func (c Circle) area() float32 {
	return math.Pi * c.r * c.r
}

func (s Square) area() float32 {
	return s.d * s.d
}

func areas(shapes ...Shape) float32 {
	result := float32(0)
	for _, shape := range shapes {
		result += shape.area()
	}
	return result
}

func (p *Person) Talk() {
	fmt.Println("Hi, I'm", p.name)
}