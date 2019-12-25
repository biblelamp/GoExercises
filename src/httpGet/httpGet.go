package main

import (
    "net/http"
    //"io/ioutil"
    "encoding/xml"
    "golang.org/x/text/encoding/charmap" // go get golang.org/x/text
    "io"
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
    Value    string   `xml:"Value"`
}

func main() {
    response, err := http.Get("http://www.cbr.ru/scripts/XML_daily.asp?date_req=20/12/2019")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    //byteValue, err := ioutil.ReadAll(response.Body)
    //if err != nil {
    //    log.Fatal(err)
    //}

    //var valutes Valutes
    //err = xml.Unmarshal(byteValue, &valutes)
    //if err != nil {
    //    log.Fatal(err)
    //}

    d := xml.NewDecoder(response.Body)
    d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
        switch charset {
        case "windows-1251":
            return charmap.Windows1251.NewDecoder().Reader(input), nil
        default:
            return nil, fmt.Errorf("Unknown charset: %s", charset)
        }
    }

    var valutes Valutes
    err = d.Decode(&valutes)
    if err != nil {
        log.Fatal(err)
    }

    //fmt.Println(valutes)

    for i := 0; i < len(valutes.Valutes); i++ {
        fmt.Println("CharCode:", valutes.Valutes[i].CharCode)
        fmt.Println("Name:", valutes.Valutes[i].Name)
        fmt.Println("Value:", valutes.Valutes[i].Value)
    }
}
