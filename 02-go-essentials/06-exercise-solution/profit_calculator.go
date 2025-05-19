package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello, World!")
	var revenvue float64
	var expenses float64
	taxRate := 0.3

	fmt.Print("Enter your revenue:")
	fmt.Scan(&revenvue)

	fmt.Print("Enter your expenses:")
	fmt.Scan(&expenses)

	fmt.Print("Enter your tax rate:")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenvue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	fmt.Printf("Hello, World! earningsBeforeTax is %f\n", earningsBeforeTax)
	fmt.Printf("Hello, World! profit is %f\n", profit)
	fmt.Printf("Hello, World! ratio is %f\n", ratio)

}
