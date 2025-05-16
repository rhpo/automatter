package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func DisplayTransitions(automate Automate) {
	colWidth := 10

	border := "+" + strings.Repeat("-", colWidth+2)
	for range automate.Alphabet {
		border += "+" + strings.Repeat("-", colWidth+2)
	}

	border += "+"

	color.New(color.FgYellow).Println(border)
	header := fmt.Sprintf("| %-*s ", colWidth, "S/L")
	for _, letter := range automate.Alphabet {
		header += fmt.Sprintf("| %-*s ", colWidth, letter)
	}
	header += "|"
	color.New(color.FgBlue).Println(header)
	color.New(color.FgYellow).Println(border)

	for _, state := range automate.Q {
		row := fmt.Sprintf("| %-*s ", colWidth, state)
		for _, letter := range automate.Alphabet {
			if transition, exists := automate.Transitions[Chemin{start: state, letter: letter}]; exists {
				row += fmt.Sprintf("| %-*s ", colWidth, transition)
			} else {
				row += fmt.Sprintf("| %-*s ", colWidth, " ")
			}
		}
		row += "|"
		color.New(color.FgGreen).Println(row)
		color.New(color.FgYellow).Println(border)
	}
}

func Display(automate Automate) {
	color.New(color.FgMagenta).Println("Alphabet:", automate.Alphabet)
	color.New(color.FgMagenta).Println("States (Q):", automate.Q)
	color.New(color.FgMagenta).Println("Initial State (Q0):", automate.Q0)
	color.New(color.FgMagenta).Println("Final States (F):", automate.F)
	color.New(color.FgMagenta).Println("Transitions:")

	DisplayTransitions(automate)
	println()
}
