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
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateNew(noteTitle, noteContent string) (*Note, error) {

	if noteTitle == "" || noteContent == "" {
		return nil, errors.New("noteTitle and noteContent is required")
	}

	return &Note{
		Title:     noteTitle,
		Content:   noteContent,
		CreatedAt: time.Now(),
	}, nil
}

func (n *Note) Display() {
	format := fmt.Sprintf("Title: %s\nContent: %s\nCreatedAt: %s\n", n.Title, n.Content, n.CreatedAt.Format(("2006-01-02 15:04:05")))
	fmt.Println(format)
}

func (n *Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
