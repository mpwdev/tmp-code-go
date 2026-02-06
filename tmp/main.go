package main

import (
	"fmt"
)

func main() {
	friends := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}

	neighbors := friends

	friends["Alice"] = 50
	fmt.Printf("friends: %#v\n", friends)
	fmt.Printf("neighbors: %#v\n", neighbors)

	people := make(map[string]int)

	for k, v := range friends {
		people[k] = v
	}
	fmt.Printf("people: %#v\n", people)

	people["Alice"] = 25
	fmt.Printf("friends: %#v\n", friends)
	fmt.Printf("people: %#v\n", people)

}
