package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// generation
	go func() {
		for x := 0; ; x++ {
			naturals <- x
			time.Sleep(time.Second * 1)
		}
	}()

	// squares
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// printing
	for {
		fmt.Println(<- squares)
	}
}
