package main

import (
	"flag"
	"fmt"
)

var (
	a *bool
	b *int
	c *string
)

func init() {
	a = flag.Bool("a", true, "the a value")
	b = flag.Int("b", 3, "the b value")
	c = flag.String("c", "empty", "the c value")
	flag.Parse()
}

func main() {
	fmt.Println(*a, *b, *c)
}