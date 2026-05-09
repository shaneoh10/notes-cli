package main

import (
	"fmt"

	"github.com/shaneoh10/notes-cli/note"
)

func main() {
	title, content, err := note.GetNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Note created successfully!\nTitle: %s\nContent: %s\n", title, content)
}
