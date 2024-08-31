package main

import "fmt"

func main() {
	usernames := make([]string, 2, 5)
	usernames[0] = "Alex"
	usernames[1] = "Lana"
	usernames = append(usernames, "George")
	usernames = append(usernames, "Rick")
	usernames = append(usernames, "Sarah") //cap is still 5 with 5 items in the slice
	usernames = append(usernames, "Chad") // here the capacity goes to 10, double the items

	fmt.Println(cap(usernames))

	for index,value:= range usernames{
		fmt.Println(index)
		fmt.Println(value)
	}
}