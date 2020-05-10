package main

import (
	"calculations/calculate"
	"fmt"
	"util"
)

func main() {
	expression := "3^2+10.5+5"

	str := "\"    \"%3,L,\"           \"FITR(A),\"  \"%4,5280*(A-FITR(A))"

	fmt.Println(util.SplitString(str))
	fmt.Println(util.SplitString("1,2, 3,4, 5"))

	//fmt.Println(calculate.Calculate(expression))

	stack := calculate.ToPostfix(expression)
	for stack.Len() > 0 {
		fmt.Print(stack.Pop(), " ")
	}
}
