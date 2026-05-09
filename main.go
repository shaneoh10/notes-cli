package main

import (
	"fmt"

	"github.com/shaneoh10/notes-cli/note"
)

func main() {
	title, content := note.GetNoteData()

	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Failed to create note:", err)
		return
	}

	newNote.Display()
	err = newNote.Save()

	if err != nil {
		fmt.Println("Failed to save note:", err)
		return
	}

	fmt.Println("Note saved successfully!")
}
