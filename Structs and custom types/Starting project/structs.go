package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	var user = User{
		// firstName: firstName,
		// lastName: lastName,
		// birthdate: birthdate,
		// createdAt: time.Now(),
		firstName,
		lastName,
		birthdate,
		time.Now(),	
	}
	outputUserDetails(user)
}
func outputUserDetails(user User){
	fmt.Print(user.firstName,user.lastName,user.birthdate,user.createdAt)
}
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
