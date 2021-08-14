package main

import (
	"fmt"
	"reflect"
)

type Cat struct {}

type Dog struct {}

type Bird struct {}

func (c Cat) voice() string {
	return "meow"
}

func (d Dog) voice() string {
	return "gaf-gaf"
}

func (b Bird) voice() string {
	return "fiu-fiu"
}

type Voice interface {
	voice() string
}

func voiceAnimals(animals ...Voice) {
	for _, a := range animals {
		fmt.Println(reflect.TypeOf(a).Name(), ":", a.voice())
	}
}

func main() {
	c := Cat{}
	d := Dog{}
	b := Bird{}
	voiceAnimals(c, d, b)
}
