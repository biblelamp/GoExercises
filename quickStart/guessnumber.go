package main

import "fmt"
import "math/rand"

func main() {
    limit := 3
    count := 0
    guess := -1
    number := rand.Intn(10)
    for count < limit && guess != number {
        fmt.Print("Guess [", limit - count, " attempts] the number (0..9): ")
        fmt.Scanln(&guess)
        if number != guess {
            if guess > number {
                fmt.Println("Your number is greater")
            } else {
                fmt.Println("Your number is less")
            }
        count++
        }
    }
    if guess == number {
        fmt.Println("You WON!");
    } else {
        fmt.Println("You lose");
    }
}