package calculate

import (
	"calculations/stack"
	"fmt"
	"log"
	"strconv"
	"math"
)

func Calculate(expression string) float64 {
	data := ToPostfix(expression)
	result := new(stack.Stack)
	for data.Len() > 0 {
		item := data.Pop()
		switch fmt.Sprint(item) {
		case "+":
			result.Push(result.Pop().(float64) + result.Pop().(float64))
		case "*":
			result.Push(result.Pop().(float64) * result.Pop().(float64))
		case "-":
			second := result.Pop().(float64)
			result.Push(result.Pop().(float64) - second)
		case "/":
			second := result.Pop().(float64)
			result.Push(result.Pop().(float64) / second)
		case "^":
			second := result.Pop().(float64)
			result.Push(math.Pow(result.Pop().(float64), second))
		default:
			result.Push(item)
		}
	}
	return result.Pop().(float64)
}

func ToPostfix(expression string) *stack.Stack {
	result := new(stack.Stack)
	stackOper := new(stack.Stack)
	for i := 0; i < len(expression); i++ {
		c := string(expression[i])
		if precedence(c) > 0 {
			for stackOper.Len() > 0 && precedence(stackOper.Peek().(string)) >= precedence(c) {
				result.Add(stackOper.Pop())
			}
			stackOper.Push(c)
		} else if c == ")" {
			x := stackOper.Pop().(string)
			for x != "(" {
				result.Add(x)
				x = stackOper.Pop().(string)
			}
		} else if c == "(" {
			stackOper.Push(c)
		} else {
			f, err := strconv.ParseFloat(c, 64)
			if err != nil {
				log.Fatal(err)
			}
			result.Add(f)
		}
	}
	for stackOper.Len() > 0 {
		result.Add(stackOper.Pop())
	}
	return result
}

func precedence(c string) int {
	switch c {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^":
		return 3
	}
	return -1
}
