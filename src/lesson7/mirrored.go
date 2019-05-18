package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() {
		responses <- request("asia.site.io")
	}()
	go func() {
		responses <- request("europe.site.io")
	}()
	go func() {
		responses <- request("americas.site.io")
	}()
	return <-responses // возврат самого быстрого ответа
}

func request(hostname string) string {
	response, err := http.Get(hostname)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(contents)
}
