package main

import "fmt"

func main() {

	price, inStock := 100, true
	if price > 80 {
		fmt.Println("Expensive")
	}

	//_ = inStock

	if price <= 100 && inStock == true {
		fmt.Println("You can buy it")
	}

}
