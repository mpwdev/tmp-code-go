package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println("--------------------------------")
	fmt.Println("Todo text:", todo.Text)
	fmt.Println("--------------------------------")
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		//fmt.Println("Error marshalling todo:", err)
		return err
	}

	err = os.WriteFile(fileName, json, 0644)
	if err != nil {
		//fmt.Println("Error writing file:", err)
		return err
	}

	return nil
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("content is required")
	}

	return Todo{Text: content}, nil
}
