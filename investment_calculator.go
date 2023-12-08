package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount int
	var expectedReturnRate float64
	var years int

	investmentAmount = 1000
	expectedReturnRate = 5.5
	years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1 + expectedReturnRate / 100, float64(years))
	fmt.Println(futureValue)
}