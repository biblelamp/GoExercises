package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Countdown started.")
	tick := make(<-chan time.Time)    // creating unidirectional chanel
	tick = time.Tick(1 * time.Second) // creating tick flow
	for countdown := 10; countdown > 0; countdown-- {
		moment := <-tick
		fmt.Println(moment.Format("15:04:05"), countdown)
	}
	fmt.Println("The rocket starts!")
}
