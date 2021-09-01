package main

import (
	"log"
	"os"
)

func main() {
	// create new file or clear file if exists
	file, err := os.Create("hello.txt")

	// error processing
	if err != nil {
		log.Fatal(err)
	}

	// write string to file
	var count int
	count, err = file.WriteString("1. Hello world!\n2. Hello world!")

	// error processing

	// logging
	log.Printf("write to file %d bytes", count)

	// opening an existing file
	file, err = os.Open("hello.txt")

	// error processing

	// prepare buffer for reading
	buffer := make([]byte, 128)

	// reading from file
	count, err = file.Read(buffer)

	// error processing

	// logging
	log.Printf("read from file %d bytes: %s", count, string(buffer[:count]))

	// close file at the end of all operations with him
	defer func() {
		err = file.Close()
		// an error is also possible when closing
		if err != nil {
			log.Fatal(err)
		}
	}()
}
