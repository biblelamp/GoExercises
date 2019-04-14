package main

import (
    "encoding/xml"
    "io/ioutil"
    "log"
)

type Notes struct {
    To      string `xml:"to"`
    From    string `xml:"from"`
    Heading string `xml:"heading"`
    Body    string `xml:"body"`
}

func main() {
    note := &Notes{To: "Шеф",
        From:    "Меня",
        Heading: "Извинения",
        Body:    "Простите, я слегка опоздаю",
    }

    file, err := xml.MarshalIndent(note, "", " ")
    if err != nil {
        log.Fatal(err)
    }

    err = ioutil.WriteFile("note1.xml", file, 0644)
    if err != nil {
        log.Fatal(err)
    }
}