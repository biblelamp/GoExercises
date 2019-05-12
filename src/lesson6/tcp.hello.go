package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		for {
			_, err = io.WriteString(conn, "Hello tcp-world!\n\r")
			if err != nil {
				log.Print(err)
				return
			}
			time.Sleep(1 * time.Second)
		}
		defer conn.Close()
	}
}
