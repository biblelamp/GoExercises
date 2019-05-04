package main

/*
 * Syntax Go. Homework 2
 * Sergey Iryupin, dated Apr 27, 2019
 */

import "fmt"

func main() {

	fmt.Printf("%d is even : %t\n", 2, isEven(2))
	fmt.Printf("%d is even : %t\n", 3, isEven(3))

	fmt.Printf("%d divided by %d : %t\n", 6, 3, isWithoutRemainder(6, 3))
	fmt.Printf("%d divided by %d : %t\n", 13, 3, isWithoutRemainder(13, 3))

	printNFibonacci(0)
	printNFibonacci(1)
	printNFibonacci(15)
	fmt.Println(sliceNFibonacci(0))
	fmt.Println(sliceNFibonacci(1))
	fmt.Println(sliceNFibonacci(15))

}

func isEven(number int) bool {
	return isWithoutRemainder(number, 2)
}

func isWithoutRemainder(number, divider int) bool {
	return number%divider == 0
}

func printNFibonacci(count int) {
	if count > 94 {
		fmt.Print("Too many Fibonacci numbers for uint64 type")
		return
	}
	if count > 0 {
		fmt.Print(0)
	} else {
		return
	}
	if count > 1 {
		fmt.Printf(" %d", 1)
	}
	var number1 uint64 = 0
	var number2 uint64 = 1
	for i := 2; i < count; i++ {
		number1, number2 = number2, (number1 + number2)
		fmt.Printf(" %d", number2)
	}
	fmt.Println()
}

func sliceNFibonacci(count int) []uint {
	if count > 94 {
		fmt.Print("Too many Fibonacci numbers for uint64 type")
		return make([]uint, 0)
	}
	fibonacci := make([]uint, count, count)
	for i := 0; i < count; i++ {
		if i < 2 {
			fibonacci[i] = uint(i)
		} else {
			fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
		}
	}
	return fibonacci
}
