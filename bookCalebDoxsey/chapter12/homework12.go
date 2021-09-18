package main

import (
	"fmt"
	"golang-book/math"
)

func main() {
	fmt.Println(math.Average([]float64{1,2}))
	fmt.Println(math.Average([]float64{}))

	fmt.Println(math.Min([]float64{1,2,3,4,5,-1}))
	fmt.Println(math.Min([]float64{}))
}
