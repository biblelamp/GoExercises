package stack

type item struct {
	value interface{} // value as interface type to hold any data type
	next  *item
}

type Stack struct {
	top    *item
	bottom *item
	size   int
}

func (stack *Stack) Len() int {
	return stack.size
}

func (stack *Stack) Add(value interface{}) {
	bottom := item{
		value: value,
		next:  nil,
	}
	if stack.Len() > 0 {
		stack.bottom.next = &bottom
	}
	stack.bottom = &bottom
	if stack.Len() == 0 {
		stack.top = stack.bottom
	}
	stack.size++
}

func (stack *Stack) Push(value interface{}) {
	stack.top = &item{
		value: value,
		next:  stack.top,
	}
	if stack.Len() == 0 {
		stack.bottom = stack.top
	}
	stack.size++
}

func (stack *Stack) Peek() (value interface{}) {
	if stack.Len() > 0 {
		value = stack.top.value
		return
	}
	return nil
}

func (stack *Stack) Pop() (value interface{}) {
	if stack.Len() > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return
	}
	return nil
}
