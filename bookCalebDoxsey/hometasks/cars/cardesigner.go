package main

import (
	"fmt"
	"reflect"
)

/*
 * Car designer. Describe Wheels, Engine and Body. Using
 * these elements "collect" several cars of different types and display
 * their constructive composition. Use structures, methods, interface.
 *
 * Автоконструктор. Описать Колеса, Двигатель и Корпус. Используя
 * эти элементы "собрать" несколько авто разных видов и отобразить
 * их конструктивный состав. Использовать структуры, методы, интерфейс.
 */

type Tracks struct {
}

type Wheels struct {
	number int
	season string
}

type Engine struct {
	cylinders int
	fuel string
}

type Body struct {
	typeBody string
}

type Car struct {
	wheels Wheels
	engine Engine
	body Body
}

type Tractor struct {
	tracks Tracks
	engine Engine
	body Body
}

func (c Car) components() {
	fmt.Println(reflect.TypeOf(c).Name(), c)
}

func (t Tractor) components() {
	fmt.Println(reflect.TypeOf(t).Name(), t)
}

func main() {

	w := Wheels{4, "summer"}
	e := Engine{4, "petrol"}
	b := Body{"sedan"}
	car := Car{w, e, b}
	car.components()

	t := Tracks{}
	tractor := Tractor{t, e, b}
	tractor.components()

}
