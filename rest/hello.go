package main

import (
	"fmt"
	"log"
	"net/http"
)

// Серия Golang: получить параметры GET и POST
// from https://russianblogs.com/article/9705261032/

// Безопасность наглядно: CORS
// and https://nuancesprog.ru/p/9316/

// HTTP заголовки по умолчанию от веб-сервера
// and https://golangify.com/customizing-http-headers#system-generated-headers

// Базовый веб сервер на Go
// and https://dev-gang.ru/article/go-web-server/

// Выразительный JavaScript: HTTP
// and https://habr.com/ru/post/245145/

// Pure and Simple - Tic Tac Toe with Javascript
// from https://dev.to/bornasepic/pure-and-simple-tic-tac-toe-with-javascript-4pgn

func helloHandler(w http.ResponseWriter, request *http.Request) {
	//io.WriteString(w, "Hello, world!\n")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "<h1>Hello</h1><p>world</p>")
	println("Hello, world!")
}

func main() {

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server running at port 8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
