package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, World!")
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Printf("Hello, World! futureValue is %f\n", futureValue)

}
