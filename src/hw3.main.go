package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
 * Syntax Go. Homework 3
 * Sergey Iryupin, dated Apr 29, 2019
 */

// Auto structure for task #1 and #2
type Auto struct {
	isPassengerCar  bool
	Brand           string
	Year            int
	Capacity        int // in litres
	isEngineRunning bool
	isWindowsOpen   bool
	CapacityFilled  int // in percent
}

// Queue for task #3
var Queue []string

// Person for phone book for task #4
type Person struct {
	Name    string
	Phone   string
	Address string
}

// PhoneBook for task #4
type PhoneBook struct {
	Persons []Person
}

func insertQueue(str ...string) {
	Queue = append(Queue, str...)
}

func removeQueue() string {
	if isEmptyQueue() {
		return nil
	}
	firstItem := Queue[0]
	Queue = Queue[1:]
	return firstItem
}

func isEmptyQueue() bool {
    return len(Queue) == 0
}

func main() {

	ford := Auto{true, "Ford Focus", 2008, 282, false, false, 0}
	tatra := Auto{false, "Tatra 815", 1983, 8000, false, false, 0}

	ford.isEngineRunning = true // turn on the engine
	tatra.CapacityFilled = 80   // place cargo

	fmt.Println(ford)
	fmt.Println(tatra)

	fmt.Println("Queue is empty: ", isEmptyQueue())
	insertQueue("first")
	insertQueue("second", "third", "fourth", "fifth")
	fmt.Println(Queue)

	removedItem := removeQueue()
	fmt.Println(removedItem, Queue)

	removedItem = removeQueue()
	fmt.Println(removedItem, Queue)

	john := Person{"John Willson", "+1 234 567-8-00", "1556 Broadway, New York, NY 10120, USA"}
	mark := Person{"Mark Fridman", "+1 432 123-4-56", "1125 16th Street, NW, Washington D.C., USA"}
	phoneBook := PhoneBook{[]Person{john, mark}}
	fmt.Println(phoneBook)

	jsonPhoneBook, err := json.Marshal(phoneBook)
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create("phoneBook.json")
	if err != nil {
		return
	}
	defer file.Close()

	file.Write(jsonPhoneBook)
}
