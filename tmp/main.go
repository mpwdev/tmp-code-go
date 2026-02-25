package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64 = 1000
	var years float64 = 10
	expectedReturnRate := 5.5

	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter Years: ")
	fmt.Scan(&years)
	fmt.Println("--------------------------------")

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
