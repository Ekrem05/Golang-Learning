package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"interface/note"
	"interface/todo"
)

type saver interface{
	Save(string) error
}

func main(){
	title, content,err:=getUserInput()
	if err !=nil{
		fmt.Print(err)
		return;
	}
	myNote:=note.NewNote(title,content)
	saveNoteErr:=save(myNote.Title,myNote);
	if saveNoteErr!=nil{
		fmt.Print(saveNoteErr.Error())
		return;
	}
	text,err:=getUserInputToDo()
	if err !=nil{
		fmt.Print(err)
		return;
	}

	todo:=todo.New(text)
	saveTodoErr:=save("todo",todo)
	if saveTodoErr!=nil{
		fmt.Print(saveTodoErr.Error())
		return;
	}
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
func getUserInputToDo() (string,error){

	fmt.Print("Text:")
	reader:=bufio.NewReader(os.Stdin)
	text,err:=reader.ReadString('\n')

	if err!=nil{
		return "",errors.New("Failed to read")
	}

	if text == "" {
		return "",errors.New("Inputs cannot be empty");
	}

	text = strings.TrimSuffix(text,"\n")
	text = strings.TrimSuffix(text,"\r")

	return text,nil
}
func save(title string,data saver)error{
	err:=data.Save(title);
	if err!=nil{
		return err;
	}
	return nil;
}
func add[T int|float64|string](a,b T) T{
	return a+b;
}