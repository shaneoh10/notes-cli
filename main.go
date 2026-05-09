package main

import (
	"fmt"

	"github.com/shaneoh10/notes-cli/note"
	"github.com/shaneoh10/notes-cli/todo"
)

func main() {
	title, content := note.GetNoteData()
	todoText := note.GetUserInput("Enter todo item: ")

	// Create new note
	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Failed to create note:", err)
		return
	}

	//Create new todo
	newTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Failed to create todo:", err)
		return
	}

	// Save note
	err = saveData(newNote)
	if err != nil {
		return
	}

	// Save todo
	err = saveData(newTodo)
	if err != nil {
		return
	}
}

type saver interface {
	Display()
	Save() error
}

func saveData(data saver) error {
	data.Display()

	var dataType string

	switch data.(type) {
	case todo.Todo:
		dataType = "Todo"
	default:
		dataType = "Note"
	}

	err := data.Save()
	if err != nil {
		fmt.Printf("Failed to save %s: %v\n", dataType, err)
		return err
	}

	fmt.Printf("%s saved successfully!\n", dataType)
	return nil
}
