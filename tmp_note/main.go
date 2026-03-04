package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error getting note data:", err)
		return
	}

	userNote.Display()

}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	//var value string
	//fmt.Scanln(&value)
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")
	//value = strings.TrimSpace(value)

	return value
}
