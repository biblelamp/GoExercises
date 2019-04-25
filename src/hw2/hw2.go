package main

/*
 * Syntax Go. Homework 2
 * Sergey Iryupin, dated Apr 25, 2019
 */

import (
	"fmt"
	//"math"
)

func main() {
}

func isEven(number int) bool {
    return isWithoutRemainder(number, 2)
}

func isWithoutRemainder(number, divider) int {
    return number % divider == 0
}