package main

import (
    "net/http"
    "io/ioutil"
    "log"
    "fmt"
)

func main() {

    response, err := http.Get("https://xml.meteoservice.ru/export/gismeteo/point/13941.xml")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(string(contents))

}