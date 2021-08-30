package main

import (
	"log"
	"os"
)

func main() {
	// create new file or clear file if exists
	file, err := os.Create("file.csv")

	// error processing
	if err != nil {
		log.Fatal(err)
	}

	// opening an existing file
	file, err = os.Open("file.csv")

	// error processing
	if err != nil {
		log.Fatal(err)
	}

	// close file at the end of all operations with him
	defer func() {
		err = file.Close()
		// an error is also possible when closing
		if err != nil {
			log.Fatal(err)
		}
	}()
}
