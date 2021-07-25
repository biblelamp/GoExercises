package main

import "fmt"

func main() {
	x := 5
	zero(x)
	fmt.Println(x)

	zeroPtr(&x)
	fmt.Println(x)

	xPtr := new(int)
	onePtr(xPtr)
}

func zero(x int) {
	x = 0
}

func zeroPtr(x *int) {
	*x = 0
}

func onePtr(x *int) {
	*x = 1
	fmt.Println(*x)
}
