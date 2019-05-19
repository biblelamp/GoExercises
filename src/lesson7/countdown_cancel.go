package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	cancel := make(chan int)
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		cancel <- 1
	}()

	fmt.Println("Countdown started. Hit [Enter] to cancel")
	tick := make(<-chan time.Time)
	tick = time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		select {
		case moment := <-tick:
			fmt.Println(moment.Format("15:04:05"), countdown)
		case <-cancel:
			fmt.Println("Launch canceled!")
			return
		}
	}
	fmt.Println("The rocket starts!")
}
