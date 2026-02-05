package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println

	myStr := strings.Repeat("ab", 5)
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", 2)
	p(myStr)

	myStr = strings.ReplaceAll("192.168.0.1", ".", "-")
	p(myStr)

	s := strings.Split("a,b,c", ",")
	p(s)
	fmt.Printf("%#v\n", s)

	s1 := strings.Split("Go for Golang", "")
	p(s1)
	fmt.Printf("%#v\n", s1)

	s = []string{"I", "learn", "Go", "language"}
	p(strings.Join(s, "-"))

	s2 := strings.TrimSpace("  \t\n Hello, World! \n\t\r  ")
	p(s2)

	s3 := strings.Trim("...Hello, World!!!!!???", ".!?")
	p(s3)
}
