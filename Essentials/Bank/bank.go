package main

import (
	"fmt"
	"os"
	"strconv"

	"bank.com/utils"
)
func getBalance()(balance float64){
	bytes,_:=os.ReadFile("balance")
	text:=string(bytes);
	balance,_=strconv.ParseFloat(text,64);
	return
}
func writeInFile(balance float64){
	balanceText := fmt.Sprint(balance);
	os.WriteFile("balance", []byte(balanceText), 0644)
}

func main() {
	var balance float64 = getBalance(); 
	for {
		utils.PrintMenu()
		var choice int
		fmt.Scan(&choice)
	
		if choice == 1{
			fmt.Println("Your balance is: ",balance)
		} else if choice == 2 {
			fmt.Print("Amount: ")
			var depositAmount float64;
			fmt.Scan(&depositAmount)
			if depositAmount<=0{
				fmt.Println("Invalid amount! ")
				continue;
			}
			balance += depositAmount;
			writeInFile(balance)
			fmt.Printf("Current balance: %.2f \n", balance)
		} else if choice == 3 {
			fmt.Print("Amount: ")
			var withdrawAmount float64;
			fmt.Scan(&withdrawAmount)
			if withdrawAmount<=0{
				fmt.Println("Invalid amount! ")
				continue;
			}
			if withdrawAmount > balance{
				fmt.Println("Invalid amount! You can't withdraw more than you have!")
				continue;
			}
			balance -= withdrawAmount;
			writeInFile(balance)
			fmt.Printf("Current balance: %.2f \n", balance)
		} else{
			fmt.Print("Goodbye!")
			break;
		}
	}
	

	

}