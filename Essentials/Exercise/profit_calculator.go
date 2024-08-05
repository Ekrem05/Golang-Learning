package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue,err1 := getUserInput("Revenue: ")
	expenses,err2 := getUserInput("Expenses: ")
	taxRate,err3 := getUserInput("Tax Rate: ")

	if err1 !=nil || err2 !=nil || err3 !=nil {
		panic(err1)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	writeInFile(fmt.Sprintf("Before taxes: %f \n Profit: %f \n Ratio: %f",ebt,profit,ratio))
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64,error) {
	fmt.Print(infoText)
	var userInput float64
	fmt.Scan(&userInput)
	if userInput<=0{
		return userInput,errors.New("Value must be greater than 0!")
	}
	return userInput,nil
}

func writeInFile(text string){
	os.WriteFile("finance.txt",[]byte(text),0644)
}
