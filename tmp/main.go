package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.4
	name = "Mob Phone"
	sold = false
}

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 3
	*price = 500.4
	*name = "Mob Phone"
	*sold = false
}

func main() {
	var quantity int = 5
	var price float64 = 300.4
	var name string = "Laptop"
	var sold bool = true
	fmt.Println("Before change:", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	fmt.Println("After change:", quantity, price, name, sold)
	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("After change by pointer:", quantity, price, name, sold)
}
