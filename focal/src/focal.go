package main

import (
	"calculations/calculate"
	"fmt"
)

func main() {
	expression := "3^2+10.5+5"

	fmt.Println(calculate.Calculate(expression))

	stack := calculate.ToPostfix(expression)
	for stack.Len() > 0 {
		fmt.Print(stack.Pop(), " ")
	}
}
