package main

import "fmt"

func main() {

	a := new(int)
	b := new(int)
	*a = 1
	*b = 2
	fmt.Printf("The value of the variable a: %x, b: %x \n", *a, *b)
	a = b
	fmt.Printf("The value of the variable a: %x, b: %x \n", *a, *b)
	*a = *a + 1
	*b = *b + 1
	fmt.Printf("The value of the variable a: %x, b: %x \n", *a, *b)

}
