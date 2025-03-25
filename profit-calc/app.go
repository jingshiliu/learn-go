package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: %")
	fmt.Scan(&taxRate)

	fmt.Println("Analyzing...")

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate / 100)
	ratio := ebt / profit

	fmt.Println("Earning before tax: $", ebt)
	fmt.Println("Earning after tax: $", profit)
	fmt.Println("Ratio (EBT / profit): ", ratio)
}
