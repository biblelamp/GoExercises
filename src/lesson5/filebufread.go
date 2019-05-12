package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("filebufread.go")
	if err != nil {
		fmt.Fprintln(os.Stderr, "file openning error:", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
