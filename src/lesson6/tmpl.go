package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const tmpl = "{{.Greeting}} {{.Name}}"

	data := struct {
		Greeting string
		Name     string
	}{"Hello", "Joe"}

	t := template.Must(template.New("").Parse(tmpl))

	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
