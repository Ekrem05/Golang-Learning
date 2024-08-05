package admin

import (
	"fmt"

	"example.com/structs/user"
)



type Admin struct{
	email string
	password string
	User user.User
}

func New(email,password string) Admin{
	return Admin{
		email,
		password,
		user.New("Admin","Admin","12/12/2024"),
	}
}
func (user Admin) OutputUserDetails(){
	fmt.Println(user.email,user.password)
}