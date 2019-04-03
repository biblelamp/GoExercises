package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("fileread.go")
    if err != nil {
        return
    }
    defer file.Close()

    // getting size of file
    stat, err := file.Stat()
    if err != nil {
        return
    }

    // reading file
    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)
    if err != nil {
        return
    }

    fmt.Println(string(bs))
}