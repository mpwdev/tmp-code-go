package main

import (
	"fmt"
)

func main() {
	a := map[string]string{"A": "X"}
	//b := map[string]string{"B": "Y"}
	b := map[string]string{"A": "X"}

	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)
	fmt.Println(s1)
	fmt.Println(s2)

	if s1 == s2 {
		fmt.Println("s1 and s2 are equal")
	} else {
		fmt.Println("s1 and s2 are not equal")
	}

}
