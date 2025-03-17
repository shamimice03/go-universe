package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cloudterms.net/note-app/note"
)

func main() {
	noteTitle := InputTaker("Please enter note title: ")
	noteContent := InputTaker("Please enter note content: ")

	newNote, _ := note.CreateNew(noteTitle, noteContent)

	newNote.DisplayNotes()
	newNote.SaveNote()

}

func InputTaker(infoText string) string {
	fmt.Print(infoText)
	var text string
	reader := bufio.NewReader(os.Stdin) // Use bufio.Reader to read the entire line including spaces
	text, _ = reader.ReadString('\n')   // Read until there is any newline
	text = strings.TrimSpace(text)      // Trim the newline character at the end

	return text
}
