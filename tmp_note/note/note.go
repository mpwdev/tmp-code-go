package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Println("--------------------------------")
	fmt.Println("Note title:", note.title)
	fmt.Println("Note content:", note.content)
	fmt.Println("Note created at:", note.createdAt)
	fmt.Println("--------------------------------")
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName)

	json, err := json.Marshal(note)
	if err != nil {
		//fmt.Println("Error marshalling note:", err)
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content are required")
	}

	return Note{title: title, content: content, createdAt: time.Now()}, nil
}
