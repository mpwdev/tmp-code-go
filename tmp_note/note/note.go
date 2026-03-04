package note

import (
	"errors"
	"fmt"
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

func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content are required")
	}

	return Note{title: title, content: content, createdAt: time.Now()}, nil
}
