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

//
type Product struct {
	price       float64
	productName string
}

func changeValuesByType(product Product) {
	product.price = 300
	product.productName = "Bicycle"
}

func changeValuesByTypeByPointer(product *Product) {
	product.price = 500.4
	product.productName = "Mob Phone"
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
	//
	gift := Product{price: 100, productName: "Watch"}
	fmt.Println("Before change:", gift)
	changeValuesByType(gift)
	fmt.Println("After change:", gift)
	changeValuesByTypeByPointer(&gift)
	fmt.Println("After change by pointer:", gift)
}
