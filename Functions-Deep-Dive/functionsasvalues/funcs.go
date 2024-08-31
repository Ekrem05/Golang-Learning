package functionsarevalues

import "fmt"

func main() {
	var numbers = []int{10, 20, 30, 40}
	mutateSlice(&numbers,double)
	fmt.Println(numbers)

	mutateSlice(&numbers,triple)

	fmt.Println(numbers)
}

func mutateSlice(numbers *[]int,operation func(int) int) {

	for index, val := range *numbers {
		(*numbers)[index]=operation(val);
	}
}
func double(number int) int{
	return number*2;
}
func triple(number int) int{
	return number*3;
}