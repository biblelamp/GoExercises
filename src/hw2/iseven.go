package main

import "fmt"

// this number is even or not

func main() {

    fmt.Println("4 if even? ", isEven(4));
    fmt.Println("5 if even? ", isEven(5));

}

func isEven(n int) bool {
    return n % 2 == 0
}