package main

import "fmt"

func main() {
	x := average([]float64{1, 2, 3, 4, 5})
	fmt.Println(x)

	a,b := getTwo()
	fmt.Println(a, b)

	fmt.Println(add(1,2,3))

	// function like variables
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1,2))

	// anonymous function
	msg := "Anonymous function"
	m := func() {
		fmt.Println(msg)
	}
	m()

	func() {
		fmt.Println(msg)
	}()

	// closures
	i := -1
	increment := func() int {
		i++
		return i
	}
	fmt.Println(increment(), increment(), increment())

	makeEven := makeEvenGenerator()
	fmt.Println(makeEven(), makeEven(), makeEven())

	fmt.Println(factorial(5))

	// defer - deferred call
	deferDemo()

	// panic and recover
	panicDemo()

	// homework
	fmt.Println("Home tasks:")

	numbers := []int{0, 1, 2, 3, 4}
	fmt.Println("Sum of slice: ", sum(numbers))

	result, flag := divideAndCheck(1)
	fmt.Println("1/2 =", result, flag)
	result, flag = divideAndCheck(2)
	fmt.Println("2/2 =", result, flag)

	fmt.Println(findMax(5, 4, 3, 21, 12))

	makeOdd := makeOddGenerator()
	fmt.Println(makeOdd(), makeOdd(), makeOdd(), makeOdd())

	fmt.Println("Fibonacci(15)=", fibonacci(15))
}

func average(xs []float64) float64 {
	total := 0.0
	for _, value := range xs {
		total += value
	}
	return total / float64(len(xs))
}

func getTwo() (int, int) {
	return 5, 6
}

func add(args ...int) int {
	total := 0
	for _, value := range args {
		total += value
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func deferDemo() {
	defer second()
	first()
}

func panicDemo() {
	defer func() {
		str := recover()
		fmt.Println("Recover:", str)
	}()
	panic("PANIC")
}

// homework

func sum(slice []int) int {
	result := 0
	for _, v := range slice {
		result += v
	}
	return result
}

func divideAndCheck(value int) (int, bool) {
	result := value / 2
	flag := value % 2 == 0
	return result, flag
}

func findMax(list ...int) int {
	result := list[0]
	for _, v := range list {
		if v > result {
			result = v
		}
	}
	return result
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fibonacci(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	return fibonacci(x - 1) + fibonacci(x - 2)
}