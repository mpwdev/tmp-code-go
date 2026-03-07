package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"Product 1", "Product 2", "Product 3", "Product 4"}
	prices := [4]float64{10.0, 20.0, 30.0, 40.0}

	fmt.Println("Product names: ", productNames)
	fmt.Println("Prices: ", prices)

	productNames[2] = "Product-3-updated"
	fmt.Println("Product names: ", productNames)

	fmt.Println("Prices:", prices[2])

	featuredPrices := prices[1:]
	featuredPrices[0] = 100.0
	fmt.Println("Featured prices: ", featuredPrices)

	fmt.Println("Prices: ", prices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))

}
