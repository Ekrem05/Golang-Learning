package arrays

import "fmt"
type Product struct{
	id int
	price float64
	title string
}
func main() {
	var hobbies = [3]string {"Weightlifting", "Beekeeping", "Reading"}
	fmt.Println(hobbies);
	fmt.Println("First hobby: ",hobbies[0:1])
	fmt.Println("More hobbies: ",hobbies[1:])

	var firstSlice = hobbies[0:2]
	var secondSlice = hobbies[:2]
	fmt.Println("First way of slicing: ",firstSlice)
	fmt.Println("Second way of slicing: ",secondSlice)

	var lastTwo = firstSlice[1:3]; //explicit
	fmt.Println("Last two: ",lastTwo)

	var goals =[]string {"Learn GO","Build an API in GO"}
	fmt.Println(goals)
	goals[1]="Learn pointers";
	fmt.Println(goals)

	goals = append(goals, "Make a project in GO");
	fmt.Println(goals)


	var products = []Product{ Product{ id: 1,price: 12.22, title: "Tire"}, Product{id:2,price: 2.33,title: "Napkin"}}
	fmt.Println(products)


	products = append(products,Product{id: 3,price: 199.23,title: "2 size bed"})
	fmt.Println(products)


}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.