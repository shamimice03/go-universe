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

func CreateNew(noteText string) (*Todo, error) {

	if noteText == "" {
		return nil, errors.New("noteTitle and noteContent is required")
	}

	return &Todo{
		Text: noteText,
	}, nil
}

func (n *Todo) Display() {
	format := fmt.Sprintf("Title: %s\n", n.Text)
	fmt.Println(format)
}

func (n *Todo) Save() error {
	fileName := "todo" + ".json"
	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
