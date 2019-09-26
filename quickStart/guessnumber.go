package main

import "fmt"
import "math/rand"
import "time"

func main() {
    count := 0
    guess := -1
    rand.Seed(time.Now().UnixNano())
    number := rand.Intn(10)
    for count < 3 && guess != number {
        fmt.Print("Guess [", 3 - count, " attempts] the number (0..9): ")
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