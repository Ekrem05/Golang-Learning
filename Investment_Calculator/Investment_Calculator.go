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
	customPrint("Investment amount:")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate:")
	fmt.Scan(&expectedReturn)
	
	fmt.Print("Years:")
	fmt.Scan(&years)

	futureValue := investmentAmount* math.Pow(1+expectedReturn/100,years)
	futureRealValue := futureValue/ math.Pow(1+inflationRate/100,years)
	foramttedFv:=fmt.Sprintf("Future value without inflation: %.1f \n",futureValue)
	fmt.Print(foramttedFv)
	fmt.Printf("Future value with inflation: %.2f ",futureRealValue)

}

func customPrint(text string){
	fmt.Println(text)
}