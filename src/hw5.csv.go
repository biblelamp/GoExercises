package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"John Wick", "director", "15000"},
		{"Mark Felton", "manager", "12000"},
		{"Luke Perry", "teacher", "8000"},
	}

	file, err := os.Create("file.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	write := csv.NewWriter(file)
	write.WriteAll(records)

	csvfile, err := os.Open("file.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()

	text, err := ioutil.ReadFile("file.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(text))

	reader := csv.NewReader(csvfile)
	records, err = reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(records)
}
