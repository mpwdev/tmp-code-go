package main

import (
	"fmt"
)

func main() {
	s1 := "Hello Go"
	fmt.Println(s1)

	fmt.Println(`He says: "Hello Go"`)

	s2 := `I like Go` //raw string
	fmt.Println(s2)

	var s3 = "I love" + " Go" // string concatenation
	fmt.Println(s3)

	fmt.Printf("s3: %s\n", s3)
	fmt.Printf("s3: %q\n", s3)

}
