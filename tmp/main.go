package main

import (
	"fmt"
)

func main() {
	var employes map[string]string
	fmt.Printf("%#v\n", employes)
	fmt.Printf("no of pairs: %d\n", len(employes))

	people := map[string]float64{}
	people["John"] = 25.5
	people["Jane"] = 30.0
	people["Jim"] = 35.5

	fmt.Printf("%#v\n", people)
	fmt.Printf("no of pairs: %d\n", len(people))

	map1 := make(map[int]int)
	map1[1] = 10
	map1[2] = 20
	map1[3] = 30
	fmt.Printf("%#v\n", map1)
	fmt.Printf("no of pairs: %d\n", len(map1))

	balances := map[string]float64{
		"USD": 1000.0,
		"EUR": 2000.0,
		"JPY": 3000.0,
	}
	fmt.Printf("%#v\n", balances)
	fmt.Printf("no of pairs: %d\n", len(balances))

	balances["USD"] = 500.0
	balances["GBP"] = 255.0
	fmt.Printf("%#v\n", balances)
	fmt.Printf("no of pairs: %d\n", len(balances))

	//balances["RON"] = 331.0

	v, ok := balances["RON"]
	if ok {
		fmt.Println("RON is in the map", v)
	} else {
		fmt.Println("RON is not in the map")
	}

	for k, v := range balances {
		fmt.Printf("key: %s, value: %#v\n", k, v)
	}

	delete(balances, "USD")
	fmt.Printf("%#v\n", balances)
	fmt.Printf("no of pairs: %d\n", len(balances))

}
