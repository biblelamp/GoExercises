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
        if guess != number {
            fmt.Println("Your number is", OneFromTwo(guess > number, "greater", "less"))
            count++
        }
    }
    fmt.Println(OneFromTwo(guess == number, "You WON!", "You lose"))
}

func OneFromTwo(cond bool, one string, two string) string {
    if cond {
        return one
    } else {
        return two
    }
}