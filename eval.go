package main

import "slices"

func checkSinglePath(a *Automate, currentState Letter, letter Letter) (bool, Letter) {
	for chemin, nextState := range a.Transitions {
		if chemin.start == currentState && chemin.letter == letter {
			return true, nextState
		}
	}

	return false, ""
}

func (a *Automate) IsAccepted(word string) bool {
	currentStates := []Letter{a.Q0}

	for _, char := range word {
		letter := Letter(char)
		nextStates := []Letter{}

		for i := 0; i < len(currentStates); i++ {
			exists, nextState := checkSinglePath(a, currentStates[i], letter)

			if exists {
				nextStates = append(nextStates, nextState)
			}
		}

		if len(nextStates) == 0 {
			return false
		}

		currentStates = nextStates
	}

	return slices.ContainsFunc(currentStates, a.IsFinalState)
}
