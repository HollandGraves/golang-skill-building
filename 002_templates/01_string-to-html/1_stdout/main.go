package main

import (
	"fmt"
)

func main() {
	name := "Holland Graves"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8" />
	<title="Hello World!" />
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`

	fmt.Println(tpl)
}
