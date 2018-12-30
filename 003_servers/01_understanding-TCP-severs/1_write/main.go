package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(conn, "\nI am writing to the sever through this connection established!\n(through the conn data type)\n")
		fmt.Fprintln(conn, "Hello World!")
		fmt.Fprintf(conn, "%v", "Hi there!")

		conn.Close()
	}
}
