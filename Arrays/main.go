package main

import "fmt"

func main() {
	var products [4]string = [4]string {"Book", "Ball","","Car"}
	products[2]="Toy"
	
	bestProducts := products[1:2]
	otherProducts:=bestProducts[:3]	
	fmt.Print(bestProducts)
	fmt.Println(otherProducts)
	fmt.Print(len(bestProducts), cap(bestProducts))

	prices := [4]int{10,30,12,5}
	fmt.Print(prices)
}