package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

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

	fmt.Print("Title:")
	reader:=bufio.NewReader(os.Stdin)
	title,err:=reader.ReadString('\n')

	if err!=nil{
		return "","",errors.New("Failed to read")
	}

	if title == "" {
		return "","",errors.New("Inputs cannot be empty");
	}

	fmt.Print("Content:")
	content,err:=reader.ReadString('\n')

	if err!=nil{
		return "","",errors.New("Failed to read")
	}

	if content == "" {
		return "","",errors.New("Inputs cannot be empty");
	}

	title = strings.TrimSuffix(title,"\n")
	title = strings.TrimSuffix(title,"\r")

	content = strings.TrimSuffix(content,"\n")
	content = strings.TrimSuffix(content,"\r")

	return title,content,nil
}