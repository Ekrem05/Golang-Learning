package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NewNote(title, content string) Note {
	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}
}
func (note Note) Save(fileName string) error {
	file,err:=os.Create(fileName+".json");
	if err!=nil{
		return errors.New("Error while creating a todo")
	}
	encoder :=json.NewEncoder(file)

	encodingErr:=encoder.Encode(note);
	if encodingErr != nil {
		return errors.New("Error encoding JSON to file")
    }
	fmt.Println("JSON data written to ", fileName)
	return nil
}