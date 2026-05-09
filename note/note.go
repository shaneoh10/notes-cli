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
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input")
	}

	return Note{
		Title:     title,
		Note:      content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func (n Note) Display() {
	fmt.Printf("Title: %s\n\nContent: %s\n\nCreated At: %s\n\n", n.Title, n.Note, n.CreatedAt.Format(time.RFC1123))
}
