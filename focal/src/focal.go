package main

import (
	"calculations/stack"
	"fmt"
)

func main() {
	stack := new(stack.Stack)

	// Push different data type to the stack
	stack.Push(1)
	stack.Push("Hello, Focal!")
	stack.Push(4.0)

	// Pop until stack is empty
	for stack.Len() > 0 {
		fmt.Println(stack.Pop())
	}
}
