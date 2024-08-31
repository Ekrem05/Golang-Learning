package note

import (
	"encoding/json"
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
func (note Note) SaveToJson(fileName string) {
	file,err:=os.Create(fileName);
	if err!=nil{
		fmt.Print("Error creating a file")
	}
	encoder :=json.NewEncoder(file)

	encodingErr:=encoder.Encode(note);
	if encodingErr != nil {
        fmt.Println("Error encoding JSON to file:", encodingErr)
        return;
    }
	fmt.Println("JSON data written to person.json")
}