package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cloudterms.net/interface/note"
	"cloudterms.net/interface/todo"
)

type saver interface {
	Save() error
}

func saveData(data saver) {
	err := data.Save()
	if err != nil {
		fmt.Println("Failed to save")
		panic(err)
	}
}

func main() {
	// note
	noteTitle := InputTaker("Please enter note title: ")
	noteContent := InputTaker("Please enter note content: ")

	newNote, _ := note.CreateNew(noteTitle, noteContent)
	newNote.DisplayNotes()
	saveData(newNote)

	// todo
	todoText := InputTaker("Enter you todo: ")

	newTodo, _ := todo.CreateNew(todoText)
	newTodo.DisplayTodos()
	saveData(newTodo)
}

func InputTaker(infoText string) string {
	fmt.Print(infoText)
	var text string
	reader := bufio.NewReader(os.Stdin) // Use bufio.Reader to read the entire line including spaces
	text, _ = reader.ReadString('\n')   // Read until there is any newline
	text = strings.TrimSpace(text)      // Trim the newline character at the end

	return text
}
