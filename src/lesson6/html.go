package main

import (
	"fmt"
	"html"
)

func main() {
	s1 := `"Fran & Freddie's Diner" <tasty@example.com>`
	s2 := html.EscapeString(s1)
	s3 := html.UnescapeString(s2)
	fmt.Println(s2)
	fmt.Println(s3)
}
