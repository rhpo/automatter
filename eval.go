package main

import "slices"

// checkSinglePath checks if there is a valid transition for the given current state and letter.
func checkSinglePath(a *Automate, currentState Letter, letter Letter) (bool, Letter) {
	for chemin, nextState := range a.Transitions {
		// trouver le chemin li fih (currentState, letter), ida kayen n9ollo kayen w nmddo win raho rayeh
		// w ida makach n9ollolo false
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
			return false // nooo valid transitions for the current letter
		}

		currentStates = nextStates // movee to the next set of states
	}

	// check if any of the current states is a final state
	return slices.ContainsFunc(currentStates, a.IsFinalState) // No final state reached
}
