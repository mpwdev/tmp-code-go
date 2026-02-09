package main

import "fmt"

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}

func main() {
	prices := []int{100, 200, 300}
	changeSlice(prices)
	fmt.Println(prices)

	pricesMap := map[string]int{"a": 100, "b": 200, "c": 300}
	changeMap(pricesMap)
	fmt.Println(pricesMap)

}
