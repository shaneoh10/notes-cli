package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	note      string
	createdAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input")
	}

	return Note{
		title:     title,
		note:      content,
		createdAt: time.Now(),
	}, nil
}

func (n Note) Save() {

}

func (n Note) Display() {
	fmt.Printf("Title: %s\n\nContent: %s\n\nCreated At: %s\n\n", n.title, n.note, n.createdAt.Format(time.RFC1123))
}
