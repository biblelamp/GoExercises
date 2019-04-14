package main

import (
    "encoding/xml"
    "io/ioutil"
    "fmt"
    "log"
)

type Notes struct {
    To      string `xml:"to"`
    From    string `xml:"from"`
    Heading string `xml:"heading"`
    Body    string `xml:"body"`
}

func main() {
    data, err := ioutil.ReadFile("note.xml")
    if err != nil {
        log.Fatal(err)
    }

    note := &Notes{}

    err = xml.Unmarshal([]byte(data), &note)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(note.To)
    fmt.Println(note.From)
    fmt.Println(note.Heading)
    fmt.Println(note.Body)
}