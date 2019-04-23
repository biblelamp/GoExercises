package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type valutes struct {
	NumCode  string  `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Nominal  int     `xml:"Nominal"`
	Name     string  `xml:"Name"`
	Value    float64 `xml:"Value"`
}

func main() {
	response, err := http.Get("http://www.cbr.ru/scripts/XML_daily.asp?date_req=20/04/2019")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//valute := &valutes{}
	/*
		err = xml.Unmarshal([]byte(contents), &valute)
		if err != nil {
			log.Fatal(err)
		}
	*/
	fmt.Print(string(contents))
}
