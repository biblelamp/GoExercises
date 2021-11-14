package main

import (
	"fmt"
	"strings"
)

func main() {
	const literal = `⌘`
	fmt.Printf("%s\n", literal)
	fmt.Printf("%+q\n", literal)
	fmt.Printf("% x\n", literal)

	fmt.Println(
		strings.Contains("test", "es"),
		// Output: true
		strings.Count("test", "t"),
		// Output: 2
		strings.HasPrefix("test", "te"),
		// Output: true
		strings.HasSuffix("test", "st"),
		// Output: true
		strings.Index("test", "e"),
		// Output: 1
		strings.Join([]string{"a","b"}, "-"),
		// Output: "a-b"
		strings.Repeat("a", 5),
		// Output: "aaaaa"
		strings.Replace("aaaa", "a", "b", 2),
		// Output: "bbaa"
		strings.Split("a-b-c-d-e", "-"),
		// Output: []string{"a","b","c","d","e"}
		strings.ToLower("TEST"),
		// Output: "test"
		strings.ToUpper("test"),
		// Output: "TEST"
	)
}
