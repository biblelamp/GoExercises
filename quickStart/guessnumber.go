package main

import "fmt"
import "math/rand"
import "time"

func main() {
    rand.Seed(time.Now().UnixNano())
    count, guess, number := 0, -1, rand.Intn(10)
    //count := 0
    //guess := -1
    //number := rand.Intn(10)
    for count < 3 && guess != number {
        fmt.Printf("Guess [%d attempts] the number (0..9): ", 3 - count)
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