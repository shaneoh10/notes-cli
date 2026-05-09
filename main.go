package main

import (
	"fmt"

	"github.com/shaneoh10/notes-cli/note"
)

func main() {
	title, content := note.GetNoteData()

	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	newNote.Display()
}
