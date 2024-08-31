package main

import "fmt"

func main() {
	var numbers = []int{10, 20, 30, 40}
	sum:=sum(numbers...)
	fmt.Println(sum)
}

func sum(numbers ...int) int {
	sum:=1
	for _, val := range numbers {
		sum+=val;
	}
	return sum;
}