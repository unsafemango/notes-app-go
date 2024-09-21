package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"unsafemango.com/note-taking-app/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	userNote.Display()

	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed!")
		return
	}

	fmt.Println("Saving the note succeeded!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

// Utility function to get user input
func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin) // stdin is the command line
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n") // exposes some utility functions for working with strings
	text = strings.TrimSuffix(text, "\r")

	// scan is meant to be used for single unit inputs so we are unable to enter longer texts (multiple spaced words would not work with scan())
	// var value string
	// fmt.Scanln(&value)

	return text
}
