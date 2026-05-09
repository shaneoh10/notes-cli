package main

import (
	"fmt"

	"github.com/shaneoh10/notes-cli/note"
	"github.com/shaneoh10/notes-cli/todo"
)

func main() {
	title, content := note.GetNoteData()
	todoText := note.GetUserInput("Enter todo item: ")

	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Failed to create note:", err)
		return
	}

	newTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Failed to create todo:", err)
		return
	}
	newNote.Display()
	err = newNote.Save()

	if err != nil {
		fmt.Println("Failed to save note:", err)
		return
	}

	fmt.Println("Note saved successfully!")

	newTodo.Display()
	err = newTodo.Save()

	if err != nil {
		fmt.Println("Failed to save todo:", err)
		return
	}

	fmt.Println("Todo saved successfully!")
}
