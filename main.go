package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {

	automate := NewAutomate()

	ReadAutomate(automate)

	Display(*automate)

	for {
		var word string
		color.New(color.FgCyan).Print("Enter a word (Ctrl-C to stop): ")
		fmt.Scanf("%s", &word)

		if automate.IsAccepted(word) {
			color.New(color.FgGreen).Println("Accepted, d'accord?")
		} else {
			color.New(color.FgRed).Println("Not accepted, c'est claiiir?")
		}
	}
}
