package main

import (
    "log"
    "os"
)

func main() {
    f, err := os.OpenFile("file.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 644)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    log.SetOutput(f)

    log.Print("save string to the log file") 
}