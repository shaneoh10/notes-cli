package note

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func GetNoteData() (string, string) {
	title := GetUserInput("Enter note title: ")
	content := GetUserInput("Enter note content: ")

	return title, content
}
