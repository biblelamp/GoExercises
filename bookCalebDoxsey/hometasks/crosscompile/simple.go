package main

// compile:
// go tool compile -N main.go

// see asm code:
// go tool objdump -S main.o

// link:
// go tool link main.o

// see support platforms:
// go tool dist list

// see target os and arch
// go env GOOS GOARCH

// example of conditional build:
// set GOOS=linux
// go build main.go

import (
	"fmt"
	"os"
)

func main() {
	var a = 1 + 1
	_, _ = fmt.Fprintln(os.Stdout, a)
}