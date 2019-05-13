package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	bs, err := ioutil.ReadFile("hw5.fileread.go")
	if err != nil {
		log.Fatal(err) // return
	}

	fmt.Println(string(bs))

}
