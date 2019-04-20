package main

import (
	"fmt"
	"math"
)

// calculation of a right triangle

func triangle() {

	var a, b float64

	fmt.Println("Enter legs of right triangle (a, b):")
	fmt.Scanln(&a, &b)

	s := a * b / 2
	g := math.Sqrt(a*a + b*b)
	p := a + b + g
	fmt.Println("s =", s, "g =", g, "p =", p)

}
