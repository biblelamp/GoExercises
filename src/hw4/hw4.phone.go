package main

import (
	"fmt"
	"sort"
)

/*
 * Syntax Go. Homework 4
 * Sergey Iryupin, dated May 04, 2019
 */

// Item for PhoneList
type Item struct {
	Name  string
	Phone string
}

// PhoneList for sorting
type PhoneList []Item

// Len is the number of elements in the collection
func (pb PhoneList) Len() int {
	return len(pb)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (pb PhoneList) Less(i, j int) bool {
	return pb[i].Name < pb[j].Name
}

// Swap swaps the elements with indexes i and j
func (pb PhoneList) Swap(i, j int) {
	pb[i], pb[j] = pb[j], pb[i]
}

func main() {
	people := []Item{
		{"Luke", "365-123-321"},
		{"Mark", "161-455-677"},
		{"John", "310-001-002"},
	}
	fmt.Println(people)

	sort.Sort(PhoneList(people))
	fmt.Println(people)
}
