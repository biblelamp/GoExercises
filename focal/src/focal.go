package main

import (
	"calculations/calculate"
	"fmt"
)

func main() {
	//stack := new(stack.Stack)

	// Push different data type to the stack
	//stack.Push(1)
	//stack.Push("Hello, Focal!")
	//stack.Push(4.5)
	//stack.Add(1)
	//stack.Add(2)
	//stack.Add("+")
	//stack.Push(1)
	//stack.Push(2)

	stack := calculate.ToPostfix("2*(1+1)")

	fmt.Println(calculate.Calculate("2*(1+1)+5"))

	//letters := []interface{}{"a", 123, "c", "d"}
	//fmt.Println(letters)

	// Pop until stack is empty
	for stack.Len() > 0 {
		fmt.Print(stack.Pop(), " ")
	}
}
