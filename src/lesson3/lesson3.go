package main

import "fmt"

type phones struct {
	home   string
	mobile string
}

type contact struct {
	name string
	call phones
}

func main() {

	addressBook := make(map[string]int)
	if number, ok := addressBook["Alex"]; ok {
		fmt.Println(number)
	}

	var person = contact{
		name: "Alex",
		call: phones{"123-45-67", "+7 919 123 4545"},
	}
	fmt.Println(person)

	person.call.home = "098-12-12"

	fmt.Println(person)
}
