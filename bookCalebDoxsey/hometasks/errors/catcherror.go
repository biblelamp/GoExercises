package main

import (
	"errors"
	"fmt"
)

func main() {
	/*defer func() {
		if v := recover(); v != nil {
			// recover actions
			fmt.Println("Recovered: ", v)
		} else {
			println("Everything is OK")
		}
	}()*/
	println("we start...")
	//z := 0
	//a := 1/z
	//println(a)
	//panic(errors.New("I don't know what to do next"))
	//panic(fmt.Errorf("I don't know what to do next"))

	err := isFloat64(1024.0)
	if err != nil {
		fmt.Println(err)
	}

	err = isFloat64("foo")
	if err != nil {
		fmt.Println(err)
	}

	println("we finish.")
}

func isFloat64(n interface{}) error {
	// is n float64?
	f, ok := n.(float64)
	if ok {
		fmt.Printf("%f is float64\n", f)
		return nil
	}
	return errors.New(fmt.Sprintf("could not assert that \"%v\" is float64.\n", n))
}
