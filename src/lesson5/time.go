package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    fmt.Println(t.Format("02-01-2006"))
    fmt.Println(t.Format("02-01-2006 15:04:05"))
}