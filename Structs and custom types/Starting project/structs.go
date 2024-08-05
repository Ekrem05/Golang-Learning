package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func (user User) outputUserDetails(){
	fmt.Println(user.firstName,user.lastName,user.birthdate,user.createdAt)
}
//the reciever is just like a normal parameter
func (user *User) clearUsername(){
	user.firstName=""
	user.lastName=""
}

func newUser(firstName,lastName,birthdate string) (*User,error) {
	if firstName == "" || lastName =="" || birthdate == ""{
		return nil, errors.New("all parameters are required");
	}
	return &User{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	},nil
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	var user *User

	user,err := newUser(firstName,lastName,birthdate)
	if err !=nil{
		fmt.Println(err)
	}
	user.outputUserDetails()
	user.clearUsername()
	user.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
