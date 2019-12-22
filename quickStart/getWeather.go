package main

import (
    "net/url"
    "net/http"
    "io/ioutil"
    "encoding/xml"
    "log"
    "fmt"
)

const WEATHER_URL string = "https://xml.meteoservice.ru/export/gismeteo/point/13941.xml"

type Weather struct {
    XMLName xml.Name `xml:"MMWEATHER"`
    Text    string   `xml:",chardata"`
    REPORT  struct {
        Text string `xml:",chardata"`
        Type string `xml:"type,attr"`
        TOWN struct {
            Text      string `xml:",chardata"`
            Index     string `xml:"index,attr"`
            Sname     string `xml:"sname,attr"`
            Latitude  string `xml:"latitude,attr"`
            Longitude string `xml:"longitude,attr"`
            FORECAST  []struct {
                Text      string `xml:",chardata"`
                Day       string `xml:"day,attr"`
                Month     string `xml:"month,attr"`
                Year      string `xml:"year,attr"`
                Hour      string `xml:"hour,attr"`
                Tod       string `xml:"tod,attr"`
                Predict   string `xml:"predict,attr"`
                Weekday   string `xml:"weekday,attr"`
                PHENOMENA struct {
                    Text          string `xml:",chardata"`
                    Cloudiness    string `xml:"cloudiness,attr"`
                    Precipitation string `xml:"precipitation,attr"`
                    Rpower        string `xml:"rpower,attr"`
                    Spower        string `xml:"spower,attr"`
                } `xml:"PHENOMENA"`
                PRESSURE struct {
                    Text string `xml:",chardata"`
                    Max  string `xml:"max,attr"`
                    Min  string `xml:"min,attr"`
                } `xml:"PRESSURE"`
                TEMPERATURE struct {
                    Text string `xml:",chardata"`
                    Max  string `xml:"max,attr"`
                    Min  string `xml:"min,attr"`
                } `xml:"TEMPERATURE"`
                WIND struct {
                    Text      string `xml:",chardata"`
                    Min       string `xml:"min,attr"`
                    Max       string `xml:"max,attr"`
                    Direction string `xml:"direction,attr"`
                } `xml:"WIND"`
                RELWET struct {
                    Text string `xml:",chardata"`
                    Max  string `xml:"max,attr"`
                    Min  string `xml:"min,attr"`
                } `xml:"RELWET"`
                HEAT struct {
                    Text string `xml:",chardata"`
                    Min  string `xml:"min,attr"`
                    Max  string `xml:"max,attr"`
                } `xml:"HEAT"`
            } `xml:"FORECAST"`
        } `xml:"TOWN"`
    } `xml:"REPORT"`
}

func main() {

    response, err := http.Get(WEATHER_URL)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    byteValue, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    var weater Weather
    err = xml.Unmarshal(byteValue, &weater)
    if err != nil {
        log.Fatal(err)
    }

    //fmt.Print(string(byteValue))
    fmt.Println(url.PathUnescape(weater.REPORT.TOWN.Sname))
    fmt.Println("Date:",
        weater.REPORT.TOWN.FORECAST[0].Day + "." +
        weater.REPORT.TOWN.FORECAST[0].Month + "." +
        weater.REPORT.TOWN.FORECAST[0].Year)
    for i := 0; i < len(weater.REPORT.TOWN.FORECAST); i++ {
        fmt.Println("T.Max:", weater.REPORT.TOWN.FORECAST[i].TEMPERATURE.Max)
    }
}