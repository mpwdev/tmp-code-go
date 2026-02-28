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

	// fmt.Print("Enter Investment Amount: ")
	// fmt.Scan(&investmentAmount)
	investmentAmount = getInvestmentAmount2()

	fmt.Print("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter Years: ")
	fmt.Scan(&years)
	fmt.Println("--------------------------------")

	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureValue := calcFutureValue(investmentAmount, expectedReturnRate, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Future Value: ", futureValue)
	fmt.Println("Future Value: ", fmt.Sprintf("%.2f", futureValue))
	fmt.Println("Future Real Value: ", fmt.Sprintf("%.2f", futureRealValue))

	fmt.Println("--------------------------------")
	outputText("text test func outputText")
}

func outputText(text string) {
	fmt.Print(text)
}

func getInvestmentAmount() float64 {
	fmt.Print("Enter Investment Amount: ")
	var investmentAmount float64
	fmt.Scan(&investmentAmount)
	return investmentAmount
}

func getInvestmentAmount2() (investmentAmount float64) {
	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)
	return investmentAmount
}

func calcFutureValue(investmentAmount float64, expectedReturnRate float64, years float64) float64 {
	return investmentAmount * math.Pow(1+expectedReturnRate/100, years)
}
