package main

import "fmt"

func main() {

	const (
		// iota starts from zero
		Jun = iota + 6
		Jul
		Aug
		Sept
		Oct
		Nov
		Dec
	)

	fmt.Println(Jun, Jul, Aug, Sept, Oct, Nov, Dec)

}
