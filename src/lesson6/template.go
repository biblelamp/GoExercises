package main

import (
	"html/template"
	"log"
	"os"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatal(err)
	}

	data := TodoPageData {
		PageTitle: "TODO list",
		Todos: []Todo {
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
