package main

import "fmt"

func main() {
    i := 1
    for i <= 5 {
        fmt.Print(i, " ")
        i = i + 1
    }
    fmt.Println()
    for i := 1; i <= 10; i++ {
        fmt.Print(i, " ")
    }
    fmt.Println()
    str := "Тест String"
    for i, ch := range str {
        fmt.Println("Code", ch, "of char in string No", i)
    }
}