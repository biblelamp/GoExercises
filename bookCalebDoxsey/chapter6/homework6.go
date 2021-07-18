package main

import "fmt"

func main() {

	// arrays
	array := []int{25,3,7,12,184,11,32}
	fmt.Println(array)
	for i, value := range array {
		fmt.Println(i, ":", value)
	}

	// slices make
	slice1 := make([]float64, 5)
	fmt.Println(slice1)

	slice2 := array[2:6]
	fmt.Println(slice2)

	// slice append
	{
		slice1 := []int{1, 2, 3}
		slice2 := append(slice1, 4, 5)
		fmt.Println(slice1, slice2)
	}

	// slice make
	{
		slice1 := []int{1, 2, 3}
		slice2 := make([]int, 2)
		copy(slice2, slice1)
		fmt.Println(slice1, slice2)
	}

	// maps
	mp := make(map[string]int)
	mp["key"] = 10
	mp["number"] = 25
	fmt.Println(mp)

	// homework
	fmt.Println(array[3])
	fmt.Println(make([]int, 3, 9))
	abc := [6]string{"a","b","c","d","e","f"}
	fmt.Println(abc[2:5])
	min := array[0]
	for i := 1; i < len(array); i++ {
		if min > array[i] {
			min = array[i]
		}
	}
	fmt.Println(min)

}
