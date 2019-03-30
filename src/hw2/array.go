package main

import "fmt"

// working with an array

var a [100]int

func main() {

    for i := 0; i < len(a); i++ {
        a[i] = i
    }

    fmt.Println(a);

}