package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost: 8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	dataCopy(os.Stdout, conn)
}

func dataCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
