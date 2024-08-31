package main

import "fmt"

func main() {
	var products [4]string = [4]string {"Book", "Ball"}
	products[2]="Toy"
	fmt.Print(products)
	fmt.Print(products[1:3])

	prices := [4]int{10,30,12,5}
	fmt.Print(prices)
}