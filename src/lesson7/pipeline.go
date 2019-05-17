package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// печать
	for {
		fmt.Println(<-squares)
	}
}
