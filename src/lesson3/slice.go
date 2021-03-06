package main

import "fmt"

func main() {

	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Println(c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)

		fmt.Printf(" %d", i+1)

		// if capacity is changed,
		// Go increases the array to adapt to new data
		if cap(scores) != c {
			c = cap(scores)
			fmt.Printf("\n%d\n", c)
		}
	}
}
