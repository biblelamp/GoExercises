package main

import (
    "fmt"
    "flag"
)

func main() {
    strPtr := flag.String("str", "foo", "a string")
    numPtr := flag.Int("num", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    var strVar string
    flag.StringVar(&strVar, "strVar", "bar", "a string var")

    flag.Parse()

    fmt.Println("str:", *strPtr)
    fmt.Println("num:", *numPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("strVar:", strVar)
    fmt.Println("tail:", flag.Args())
}