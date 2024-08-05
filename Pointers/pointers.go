package main

import "fmt"

func main() {
	var age int = 30
	var agePointer *int = &age;

	fmt.Printf("Age before the function: %.v \n",*agePointer)
	getAdultYears(agePointer);
	fmt.Printf("Age after the function: %.v \n",age)

}

func getAdultYears(age *int) {
	 *age=*age - 18;
}