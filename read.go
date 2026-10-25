package main

import (
	. "automate/functions"
	"fmt"

	"github.com/fatih/color"
)

func ReadAlphabet(automate *Automate) {
	var status *bool
	var message string

	for {
		var letter Letter

		fmt.Printf("Î£[%d]: ", len(automate.Alphabet)+1)
		fmt.Scanf("%s", &letter)

		status, message = automate.AddLetter(letter)

		if status == nil {
			ClearLine()
			break
		}

		if !*status {
			ClearLine()
			color.New(color.FgRed).Println("Error:", message)
			continue
		}
	}
}

func ReadStates(automate *Automate) {
	var status *bool
	var message string

	for {
		var state Letter

		fmt.Printf("Q[%d]: ", len(automate.Q)+1)
		fmt.Scanf("%s", &state)

		status, message = automate.AddState(state)

		if status == nil {
			ClearLine()

			break
		}

		if !*status {
			ClearLine()
			color.New(color.FgRed).Println("Error:", message)
			continue
		}
	}
}

func ReadInitialState(automate *Automate) {
	var status *bool
	var message string

	for {
		var state Letter

		fmt.Printf("Initial state (ex. %s) = ", automate.Q[0])
		fmt.Scanf("%s", &state)

		status, message = automate.AddInitialState(state)

		if status == nil {
			ClearLine()

			break
		}

		if !*status {
			ClearLine()
			color.New(color.FgRed).Println("Error:", message)
			continue
		}

		break
	}
}

func ReadFinalStates(automate *Automate) {
	var status *bool
	var message string

	for {
		var state Letter

		fmt.Printf("F[%d] = ", len(automate.F)+1)
		fmt.Scanf("%s", &state)

		status, message = automate.AddFinalState(state)

		if status == nil {
			ClearLine()

			break
		}

		if !*status {
			ClearLine()
			color.New(color.FgRed).Println("Error:", message)
			continue
		}
	}
}

func ReadTransitions(automate *Automate) {

	transition := Transition{}

	for {

		prompt := "Input the Transition (Initial Transition q):\n"

		for {
			var q0 Letter
			for {
				fmt.Print(prompt)
				fmt.Scanf("%s", &q0)


				if len(q0) == 0 && len(automate.Transitions) == 0 {
					ClearLine(2)
					fmt.Println("Please input at least one transition")
					continue
				}

				if len(q0) == 0 {
					ClearLine(2)
					status, message := automate.FinishedTransitions()
					if !*status {
						fmt.Println(message)
						continue
					}

					return
				}

				if !automate.IsFinalState(q0) && !automate.IsInitialState(q0) {
					ClearLine(2)

					fmt.Println("State does not exist")
					continue
				}

				transition.start = q0
				break
			}

			ClearLine(2)

			var letter Letter
			prompt = "Input the Transition (Transition letter):\n" + string(transition.start) + " with "

			for {
				fmt.Print(prompt)
				fmt.Scanf("%s", &letter)

				if !automate.IsLetter(letter) {
					ClearLine(2)

					fmt.Println("Letter does not exist")

					continue
				}

				break
			}

			ClearLine(2)

			var q1 Letter
			prompt = "Input the Transition (Final Transition q):\n" + string(transition.start) + " with " + string(letter) + ": "

			for {
				fmt.Print(prompt)
				fmt.Scanf("%s", &q1)

				if len(q1) == 0 {
					ClearLine(2)
					fmt.Printf("Please enter a final state\n")

					continue
				}

				if !automate.IsFinalState(q1) && !automate.IsInitialState(q1) {
					ClearLine(2)

					fmt.Printf("State '%s' does not exist in the states\n", q1)
					continue
				}

				transition.letter = letter
				transition.end = q1

				break
			}

			var status *bool
			var message string
			if !automate.TransitionExists(transition) {
				status, message = automate.AddTransition(&transition)

				if !*status {
					println("Error: " + message)
					continue
				}

			} else {
				println("Transition already exists... Please input another")
				continue
			}

			automate.AddTransition(&transition)

			println()
			break
		}

	}

}

func ReadAutomate(automate *Automate) {

	Clear()

	fmt.Println("Input the alphabet:")
	ReadAlphabet(automate)
	println()

	fmt.Println("Input the states:")
	ReadStates(automate)
	println()

	ReadInitialState(automate)
	println()

	fmt.Println("Input the final states:")
	ReadFinalStates(automate)
	println()

	ReadTransitions(automate)
	println()
}
