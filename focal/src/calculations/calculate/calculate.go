package calculate

import (
	"calculations/stack"
	"fmt"
	"log"
	"strconv"
	"strings"
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
	expression = strings.ReplaceAll(expression, " ", "")
	result := new(stack.Stack)
	stackOper := new(stack.Stack)
	numberOrVariable := ""
	for i := 0; i < len(expression); i++ {
		c := string(expression[i])

		// add number or variable
		if precedence(c) > -1 && len(numberOrVariable) > 0 {
			if value, err := strconv.ParseFloat(numberOrVariable, 64); err != nil {
				log.Fatal(err)
			} else {
				result.Add(value)
			}
			numberOrVariable = ""
		}

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
			numberOrVariable += c
		}
	}
	if len(numberOrVariable) > 0 {
		if value, err := strconv.ParseFloat(numberOrVariable, 64); err != nil {
			log.Fatal(err)
		} else {
			result.Add(value)
		}
	}
	for stackOper.Len() > 0 {
		result.Add(stackOper.Pop())
	}
	return result
}

func isFloat(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func precedence(c string) int {
	switch c {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^":
		return 3
	case "(", ")":
		return 0
	}
	return -1
}
