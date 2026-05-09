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

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}

func (t Todo) Save() error {
	fileName := "todo.json"

	data, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}

func (t Todo) Display() {
	fmt.Printf("\nTodo: %s\n\n", t.Text)
}
