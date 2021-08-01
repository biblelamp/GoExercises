package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(25 * time.Millisecond)
	x := 45
	fmt.Printf("Fibonacci(%d) = %d", x, fibonacci(x))
}

func spinner(duration time.Duration) {
	for {
		for _, v := range "-\\|/" {
			fmt.Printf("%c\r", v)
			time.Sleep(duration)
		}
	}
}

func fibonacci(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	return fibonacci(x - 1) + fibonacci(x - 2)
}