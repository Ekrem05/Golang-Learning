package main

import (
	"fmt"
)

func main() {
	revenue, expenses, taxRate := readAndReturn();

	
	earningsBeforeTax,profit,ratio:=calculate(revenue,expenses,taxRate);

	fmt.Println("Earnings before tax: ",earningsBeforeTax)
	fmt.Println("Profit: ",profit)
	fmt.Printf("Ratio: %.3f",ratio)

}

func readAndReturn()(revenue float64,expenses float64,taxRate float64){
	
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue);

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses);

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate);

	return
}
func calculate(revenue, expenses, taxRate float64)(earningsBeforeTax, profit, ratio float64){
	
	earningsBeforeTax=revenue-expenses;
	profit=earningsBeforeTax-(earningsBeforeTax*taxRate/100)
	ratio =earningsBeforeTax/profit

	return
}