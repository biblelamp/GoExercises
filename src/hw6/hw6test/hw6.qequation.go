package hw6test

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Println("Enter a, b, c:")
	fmt.Scanln(&a, &b, &c)
	x1, x2, ok := quadEquation(a, b, c)
	if ok {
		fmt.Println("x1 =", x1, "x2 =", x2)
	} else {
		fmt.Println("No root of equation, d < 0")
	}

}

// solution quadratic equation

func quadEquation(a, b, c float64) (x1, x2 float64, ok bool) {

	d := b*b - 4*a*c

	if d < 0 {
		ok = false
	} else {
		ok = true
		x1 = (-b + math.Sqrt(d)) / (2 * a)
		x2 = (-b - math.Sqrt(d)) / (2 * a)
	}
	return
}
