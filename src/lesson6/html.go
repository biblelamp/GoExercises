package main

import (
	"fmt"
	"html"
)

func main() {
	s1 := `"Fran & Freddie's Diner" <tasty@example.com>`
	s2 := html.EscapeString(s1)
	fmt.Println(s2)
	fmt.Println(html.UnescapeString(s2))
}
