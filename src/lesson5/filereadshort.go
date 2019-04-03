package main

import (
    "fmt"
    "io/ioutil"
)

func main() {

    bs, err := ioutil.ReadFile("filereadshort.go")
    if err != nil {
        return
    }

    fmt.Println(string(bs))

}