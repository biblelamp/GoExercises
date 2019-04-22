package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		if str != nil {
			fmt.Println(str)
		}
		fmt.Println("ok")
	}()
	b := 0
	fmt.Println(4 / b)
	//panic("Artifical error")

	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(factorial(15))
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func average(xs ...float64) (length int, total float64) {
	for _, v := range xs {
		total += v
	}
	length = len(xs)
	total = total / float64(length)
	return
}
