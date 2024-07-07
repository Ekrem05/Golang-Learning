package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Invest!!")
	const inflationRate = 2.5;
	var investmentAmount float64;
	var expectedReturn float64 ;
	var years float64;
	fmt.Print("Investment amount:")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate:")
	fmt.Scan(&expectedReturn)
	
	fmt.Print("Years:")
	fmt.Scan(&years)

	futureValue := investmentAmount* math.Pow(1+expectedReturn/100,years)
	futureRealValue := futureValue/ math.Pow(1+inflationRate/100,years)

	fmt.Println("Future value without inflation: ",futureValue)
	fmt.Println("Future value with inflation: ",futureRealValue)

}