package calculate

import "calculations/stack"

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
			result.Add(c)
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
