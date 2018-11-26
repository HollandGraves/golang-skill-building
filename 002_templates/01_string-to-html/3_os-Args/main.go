package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// os.Args takes arguments from the command line. os.Args[1] is the first value typed into the command line after main.go
	name := os.Args[1]
	// os.Args[0] is the filePath of the main.go file
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	// fmt.Sprint allows you to dynamically attatch variables to a string.
	// If the variables contain a value at any point they will be displayed
	// when the fmt method runs.
	// fmt.Sprint (i.e. fmt.S[tring]print) returns a string
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8" />
		<title>Hello World</title>
		</head>
		<body>
		<!-- name is a variable that is tied to the name variable in the main.go file -->
		<h1>` + name + `</h1>
		</body>
		</html>
	`)

	// os.Create creates a file struct key with the value of the string in its argument
	// it returns a *File and error type
	nf, err := os.Create("index.html")
	// always err check after
	if err != nil {
		// log.Fatal prints whatever is it's argument, then panics and closes the program
		log.Fatal("error creating files", err)
	}
	// defer will allow for whatever method is after it to be run as the last thing that is run in the program
	// os.Close() which has a reciever of (f *File) will not allow the *File type access to I/O
	defer nf.Close()

	/*
		io.Copy requires a Writer and a Reader type.

		A Writer type uses write(b []byte) method.
		nf is a file type, but since the file type has a method
		attached that uses the write(b []byte) function, it is also a Writer type
		(due to the Writer type being an interface)

		a Reader type uses the read(p []byte) method
		str is a string, but the strings.NewReader() method takes str and wraps
		its information in a reader type

		In sum: a Reader type reads from a data source,
		turns the data source into a slice of bytes,
		the Writer takes the slice of bytes and writes to the target file
		DATA SOURCE ==> READER ==> []BYTE ==> WRITER ==> TARGET RESOURCE

		io.Copy handles the process of Reading from a DATA SOURCE and Writing to a TARGET RESOURCE
	*/
	io.Copy(nf, strings.NewReader(str))

	/*
		at the terminal:
		go run main.go "Holland Graves"

		this will cause whatever is in the parantheses to be os.Args[1] so you can add in from the command line
	*/
}
