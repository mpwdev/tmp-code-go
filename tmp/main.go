package main

import "fmt"

func main() {

	const secPerDay = 60 * 60 * 24
	const daysYear = 365

	fmt.Println("Total number sec in year:", secPerDay*daysYear)

}
