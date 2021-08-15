package main

/*
 * Create several animals (at least 3 - cat, dog and bird).
 * Each animal "sounds" in its own way. Write a function where everything
 * animals, from the list, "give a voice". Use structures, methods, interface.
 *
 * Создать несколько животных (минимум 3 - кошка, собака и птичка).
 * Каждое животное "звучит" по-своему. Написать функцию, где все животные,
 * из списка, "подают голос". Использовать структуры, методы, интерфейс.
 */

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
