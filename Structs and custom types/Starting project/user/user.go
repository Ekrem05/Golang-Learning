package user

import (
	"fmt"
	"time"
)
type User struct {
	FirstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func (user User) OutputUserDetails(){
	fmt.Println(user.FirstName,user.lastName,user.birthdate,user.createdAt)
}
//the reciever is just like a normal parameter
func (user *User) ClearUsername(){
	user.FirstName=""
	user.lastName=""
}

func New(firstName, lastName, birthdate string) (User) {	
	return User{
		FirstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}