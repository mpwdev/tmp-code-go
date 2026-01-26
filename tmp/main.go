package main

import (
	"fmt"
	"time"
)

func main() {

	// lang := "golang"

	// switch lang {
	// case "Python":
	// 	fmt.Println("it is Python")
	// case "Go", "golang":
	// 	fmt.Println("It is Go")
	// default:
	// 	fmt.Println("Other langs")
	// }

	hour := time.Now().Hour()
	//fmt.Println(hour)

	switch true {
	case hour < 12:
		fmt.Println("AM")
	case hour > 12:
		fmt.Println("PM")
	}

}
