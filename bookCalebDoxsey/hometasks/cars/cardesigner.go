package main

import "fmt"

/*
 * Car designer. Describe Wheels, Engine and Body. Using
 * these elements "collect" several cars of different types and display
 * their constructive composition. Use structures, methods, interface.
 *
 * Автоконструктор. Описать Колеса, Двигатель и Корпус. Используя
 * эти элементы "собрать" несколько авто разных видов и отобразить
 * их конструктивный состав. Использовать структуры, методы, интерфейс.
 */

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
	carBody Body
}

func (c Car) show() {
	fmt.Println(c)
}

func main() {

	w := Wheels{4, "summer"}
	e := Engine{4, "petrol"}
	b := Body{"sedan"}
	c := Car{w, e, b}
	c.show()

}
