package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:\ncopy <source_file_name> <dest_file_name>")
		os.Exit(0)
	}

	source, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer source.Close()

	destination, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer destination.Close()

	copied, err := io.Copy(destination, source)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Copied %d bytes.", copied)
}
