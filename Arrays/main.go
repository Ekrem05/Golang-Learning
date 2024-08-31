package main

import "fmt"

func main() {
	var prices = []float64{11.2};
	prices = append(prices, 12.99);
	prices = append(prices, 30.11);
	fmt.Print(prices)
	
}