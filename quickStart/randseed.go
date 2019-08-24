package main

import "fmt"
import "math/rand"
import "time"

func main() {
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < 10; i++ {
        fmt.Println(rand.Intn(5))
    }
}