package note

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)
type Note struct {
	title     string    `json:"title"`
	content   string    `json:"content"`
	createdAt time.Time `json:"created_at"`
}

func NewNote(title, content string) Note {
	return Note{
		title: title,
		content: content,
		createdAt: time.Now(),
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