package main

import "fmt"

func main() {
	for i := 0; i < 5; i ++ {
		fmt.Println(i)
	}

	for i := 0; i < 5; i ++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	for i := 0; i < 5; i ++ {
		switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		default: fmt.Println("Unknown Number")
		}
	}

	array := []int{25,20,15,10,5}
	for i, value := range array {
		fmt.Println(i, ":", value)
	}
}
