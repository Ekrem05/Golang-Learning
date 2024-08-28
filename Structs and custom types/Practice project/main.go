package main

import (
	"errors"
	"fmt"

	"example.com/note/note"
)

func main(){
	title, content,err:=getUserInput()
	if err !=nil{
		fmt.Print(err)
		return;
	}
	myNote:=note.NewNote(title,content)
	myNote.SaveToJson("firstNote")
}
	


func getUserInput() (string,string,error){
	var title,content string
	fmt.Print("Title:")
	fmt.Scanln(&title)
	if title == "" {
		return "","",errors.New("Inputs cannot be empty");
	}
	fmt.Print("Content:")
	fmt.Scanln(&content)
	if content == "" {
		return "","",errors.New("Inputs cannot be empty");
	}
	return title,content,nil
}