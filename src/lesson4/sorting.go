package main

import (
	"fmt"
	"sort"
)

// Person for list
type Person struct {
	Name string
	Age  int
}

// List for sorting
type List []Person

// Len is the number of elements in the collection
func (list List) Len() int {
	return len(list)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (list List) Less(i, j int) bool {
	return list[i].Name < list[j].Name
}

// Swap swaps the elements with indexes i and j
func (list List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func main() {
	people := []Person{
		{"Luke", 36},
		{"Mark", 16},
		{"John", 31},
	}
	fmt.Println(people)

	sort.Sort(List(people))
	fmt.Println(people)
}
