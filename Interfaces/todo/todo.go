package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)
type Todo struct {
	Text   string    `json:"text"`
}

func New(text string) Todo {
	return Todo{
		Text: text,
	}
}
func (todo Todo) Save(fileName string) error {
	file,err:=os.Create(fileName+".json");
	if err!=nil{
		return errors.New("Error while creating a todo")
	}
	encoder :=json.NewEncoder(file)

	encodingErr:=encoder.Encode(todo);
	if encodingErr != nil {
		return errors.New("Error encoding JSON to file")
    }
	fmt.Println("JSON data written to ", fileName)
	return nil
}