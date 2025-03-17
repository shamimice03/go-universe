package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cloudterms.net/embedded-interface/mirror"
	"cloudterms.net/embedded-interface/note"
	"cloudterms.net/embedded-interface/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func saveData(data outputtable) {
	data.Display()
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
	saveData(newNote)

	// todo
	todoText := InputTaker("Enter you todo: ")
	newTodo, _ := todo.CreateNew(todoText)
	saveData(newTodo)

	// mirror
	mirrorText := InputTaker("Enter you mirror text: ")
	newMirror := mirror.CreateNew(mirrorText)
	saveData(newMirror)

}

func InputTaker(infoText string) string {
	fmt.Print(infoText)
	var text string
	reader := bufio.NewReader(os.Stdin) // Use bufio.Reader to read the entire line including spaces
	text, _ = reader.ReadString('\n')   // Read until there is any newline
	text = strings.TrimSpace(text)      // Trim the newline character at the end

	return text
}
