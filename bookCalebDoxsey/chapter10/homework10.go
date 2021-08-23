package main

import (
	"fmt"
	"time"
)

func main() {
	//go spinner(25 * time.Millisecond)
	//x := 45
	//fmt.Printf("Fibonacci(%d) = %d", x, fibonacci(x))

	c := make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	time.Sleep(time.Second * 6)
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

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}