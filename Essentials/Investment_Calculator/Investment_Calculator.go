package main

import (
	"fmt"
	"math"
)
const inflationRate = 2.5;
func main() {
	fmt.Println("Invest!!")
	
	var investmentAmount float64;
	var expectedReturn float64 ;
	var years float64;
	customPrint("Investment amount:")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate:")
	fmt.Scan(&expectedReturn)
	
	fmt.Print("Years:")
	fmt.Scan(&years)

	futureValue, futureRealValue:=calculateFutureValues(investmentAmount,expectedReturn,years)
	foramttedFv:=fmt.Sprintf("Future value without inflation: %.1f \n",futureValue)
	fmt.Print(foramttedFv)
	fmt.Printf("Future value with inflation: %.2f ",futureRealValue)

}

func customPrint(text string){
	fmt.Println(text)
}

func calculateFutureValues(investmentAmount, expectedReturn, years float64)(futureValue float64,futureRealValue float64){
	futureValue = investmentAmount* math.Pow(1+expectedReturn/100,years);
	futureRealValue = futureValue/ math.Pow(1+inflationRate/100,years);
	//return futureValue, futureRealValue;
	return;
}