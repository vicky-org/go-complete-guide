package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, World!")
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	years := 10.0

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Printf("Hello, World! futureValue is %f\n", futureValue)

}
