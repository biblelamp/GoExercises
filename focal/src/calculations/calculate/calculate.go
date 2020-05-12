package calculate

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"util/stack"
)

func Calculate(expression string, variables map[string]float64) float64 {
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
			if reflect.TypeOf(item).String() == "float64" {
				result.Push(item)
			} else {
				if value, ok := variables[item.(string)]; ok {
					result.Push(value)
				} else {
					result.Push(.0)
				}
			}
		}
	}
	return result.Pop().(float64)
}

func ToPostfix(expression string) *stack.Stack {
	expression = strings.ReplaceAll(expression, " ", "")
	result := new(stack.Stack)
	stackOper := new(stack.Stack)
	numberOrVariable := ""
	for i, item := range expression {
		c := string(item)

		// add number or variable
		if precedence(c) > -1 && len(numberOrVariable) > 0 {
			if value, err := strconv.ParseFloat(numberOrVariable, 64); err != nil {
				result.Add(numberOrVariable)
			} else {
				result.Add(value)
			}
			numberOrVariable = ""
		}

		if precedence(c) > 0 {
			// for expressions like: -3; 2*(-2-3)
			if c == "-" && (result.Len() == 0 || expression[i-1:i] == "(") {
				result.Add(.0)
			}
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
			result.Add(numberOrVariable)
		} else {
			result.Add(value)
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
	case "(", ")":
		return 0
	}
	return -1
}
