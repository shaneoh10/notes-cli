package note

import (
	"errors"
	"fmt"
)

func GetUserInput(prompt string) (string, error) {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)

	if input == "" {
		return input, errors.New("Invalid input")
	}

	return input, nil
}

func GetNoteData() (string, string, error) {
	title, err := GetUserInput("Enter note title: ")

	if err != nil {
		return "", "", err
	}

	content, err := GetUserInput("Enter note content: ")

	if err != nil {
		return "", "", err
	}

	return title, content, nil
}
