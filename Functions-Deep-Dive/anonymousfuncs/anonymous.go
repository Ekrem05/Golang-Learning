package anonumousfuncs

import "fmt"

func main() {
	var numbers = []int{10, 20, 30, 40}
	mutateSlice(&numbers, func(i int) int {return i*2})
	fmt.Println(numbers)
}

func mutateSlice(numbers *[]int,operation func(int) int) {

	for index, val := range *numbers {
		(*numbers)[index]=operation(val);
	}
}