package main

import "fmt"
import "math"

const inflationRate float64 = 2.5

// func overloading is not supported by go
func calculateNetFutureValue(principal float64, rate float64, years int) float64{
	return float64(principal) * math.Pow((1 + rate / 100), float64(years))
}

func calculateRealFutureValue(principal float64, rate float64, years int) float64{
	netValue := float64(principal) * math.Pow((1 + rate / 100), float64(years))
	return netValue / math.Pow(1 + inflationRate / 100, float64(years))
}

func calculateRealFutureValueWithNetValue(netValue float64, years int) float64{
	return netValue / math.Pow(1 + inflationRate / 100, float64(years))
}

func main() {
	// var investAmount int = 1000
	// var rate float64 = 5.5
	// var years int = 10

	var investAmount float64
	var rate float64
	var years int

	fmt.Println("What is the initial amount?")
	fmt.Scan(&investAmount)

	// cant guarantee type
	fmt.Println("What is the rate?")
	fmt.Scan(&rate)

	fmt.Println("What is the years?")
	fmt.Scan(&years)

	fmt.Println("Investment Amount: ", investAmount)
	fmt.Println("Rate: ", rate)
	fmt.Println("Years: ", years)

	var futureValue = calculateNetFutureValue(investAmount, rate, years)
	fmt.Println("Total: ", futureValue)
}
