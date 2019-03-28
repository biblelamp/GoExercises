package main

import ("fmt"; "math")

// solution quadratic equation

var a, b, c float64

func main() {

    fmt.Println("Enter a, b, c:")
    fmt.Scanln(&a, &b, &c)

    d := b*b - 4*a*c;

    //fmt.Println(d)

    if d < 0 {
        fmt.Println("No root of equation")
    } else if d == 0 {
        x := -b/(2*a)
        fmt.Println("x =", x)
    } else {
        x1 := (-b + math.Sqrt(d))/(2*a)
        x2 := (-b - math.Sqrt(d))/(2*a)
        fmt.Println("x1 =", x1, "x2 =", x2)
    }
}