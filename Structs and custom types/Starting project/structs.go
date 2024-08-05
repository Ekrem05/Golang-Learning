package main

import (
	"fmt"

	"example.com/structs/admin"
	"example.com/structs/user"
)


func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	var appUser user.User
	var adminUser admin.Admin
	appUser = user.New(firstName,lastName,birthdate)
	adminUser = admin.New("Admin@gmail.com","secretPassword")
	adminUser.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUsername()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
