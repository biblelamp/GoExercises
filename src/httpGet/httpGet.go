package main

import (
    "net/http"
    "io/ioutil"
    "encoding/xml"
    "strings"
    "log"
    "fmt"
)

type Valutes struct {
    XMLName  xml.Name `xml:"ValCurs"`
    Valutes  []Valute `xml:"Valute"`
}

type Valute struct {
    XMLName  xml.Name `xml:"Valute"`
    NumCode  string   `xml:"NumCode"`
    CharCode string   `xml:"CharCode"`
    Nominal  int      `xml:"Nominal"`
    Name     string   `xml:"Name"`
    Value    float64  `xml:"Value"`
}

func main() {
    response, err := http.Get("http://www.cbr.ru/scripts/XML_daily.asp?date_req=20/12/2019")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    byteValue, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    data := string(byteValue);
    data = strings.Replace(data, "windows-1251", "UTF-8", 1)

    //var valutes Valutes
    //err = xml.Unmarshal([]byte(data), &valutes)
    //if err != nil {
    //    log.Fatal(err)
    //}

    fmt.Println(data)

    //for i := 0; i < len(valutes.Valutes); i++ {
    //    fmt.Println("CharCode:", valutes.Valutes[i].CharCode)
    //    fmt.Println("Value:", valutes.Valutes[i].Value)
    //}
}
