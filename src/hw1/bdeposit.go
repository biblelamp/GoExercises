package main

import "fmt"

// bank deposit with annual %

var s, p float64

func main() {

    fmt.Println("Enter amount and annual % (s, p):")
    fmt.Scanln(&s, &p)

    s += s/100*p;
    s += s/100*p;
    s += s/100*p;
    s += s/100*p;
    s += s/100*p;
    fmt.Println("s =", s);
}