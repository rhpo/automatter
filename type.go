package main

import (
	"slices"
)

type Letter string

type Chemin struct {
	start  Letter
	letter Letter
}

type Automate struct {
	Transitions map[Chemin]Letter
	Alphabet    []Letter
	F           []Letter
	Q           []Letter
	Q0          Letter
}

type Transition struct {
	start  Letter
	letter Letter
	end    Letter
}

// transitions is a map of map , like... you access with a state + a letter to get a state
func NewAutomate() *Automate {
	return &Automate{
		Transitions: make(map[Chemin]Letter),
		Alphabet:    []Letter{},
		F:           []Letter{},
		Q:           []Letter{},
		Q0:          "",
	}
}

func (a *Automate) IsInitialState(state Letter) bool {
	return slices.Contains(a.Q, state)
}

func (a *Automate) IsLetter(letter Letter) bool {
	return slices.Contains(a.Alphabet, letter)
}

func (a *Automate) IsFinalState(state Letter) bool {
	return slices.Contains(a.F, state)
}

func (a *Automate) AddLetter(letter Letter) (success *bool, message string) {
	noError := true

	if len(letter) == 0 && len(a.Alphabet) == 0 {
		noError = false
		return &noError, "Please enter at least one letter"
	}

	if len(letter) == 0 {
		return nil, "End of input"
	}

	// Check if the letter is already in the alphabet
	if a.IsLetter(letter) {
		noError = false
		return &noError, "Letter already exists in the alphabet"
	}

	// add the letter to the alphabet
	a.Alphabet = append(a.Alphabet, letter)
	return &noError, "Letter added successfully"
}

// Q
func (a *Automate) AddState(state Letter) (success *bool, message string) {
	noError := true

	if len(state) == 0 && len(a.Q) == 0 {
		noError = false
		return &noError, "Please enter at least one state"
	}

	if len(state) == 0 {
		return nil, "End of input"
	}

	// Check if the state is already in the list of states
	if a.IsInitialState(state) {
		noError = false
		return &noError, "State already exists in the list of states"
	}

	// add the state to the list of states
	a.Q = append(a.Q, state)
	return &noError, "State added successfully"
}

func (a *Automate) AddInitialState(state Letter) (success *bool, message string) {
	noError := true

	if len(state) == 0 {
		noError = false
		return &noError, "Please enter a letter of the Alphabet"
	}

	// Check if the state is not in the list of states
	if !a.IsInitialState(state) {
		noError = false
		return &noError, "State does not exist in the list of states"
	}

	// add the state to the list of states
	a.Q0 = state
	return &noError, "Initial state added successfully"
}

func (a *Automate) AddFinalState(state Letter) (success *bool, message string) {
	noError := true

	if len(state) == 0 && len(a.F) == 0 {
		noError = false
		return &noError, "Please add at least one final state before proceeding"
	}

	if len(state) == 0 {
		return nil, "End of input"
	}

	// Check if the state is not in the list of states
	if !a.IsInitialState(state) {
		noError = false
		return &noError, "State does not exist in the list of states"
	}

	// Check if the state is already in the list of final states
	if a.IsFinalState(state) {
		noError = false
		return &noError, "State already exists in the list of final states"
	}

	// add the state to the list of final states
	a.F = append(a.F, state)
	return &noError, "Final state added successfully"
}

func (a *Automate) TransitionExists(t Transition) bool {
	found := false
	for chemin, end := range a.Transitions {
		if chemin.start == t.start && chemin.letter == t.letter && end == t.end {
			found = true
			break
		}
	}
	return found
}

func (a *Automate) AddTransition(t *Transition) (success *bool, message string) {
	noError := true

	// Check if the start state exists
	if !a.IsInitialState(t.start) {
		noError = false
		return &noError, "Start state does not exist in the list of states"
	}

	// Check if the end state exists
	if !a.IsInitialState(t.end) {
		noError = false
		return &noError, "End state does not exist in the list of states"
	}

	// Check if the letter exists in the alphabet
	if !a.IsLetter(t.letter) {
		noError = false
		return &noError, "Letter does not exist in the alphabet"
	}

	// Check if the transition already exists
	if a.TransitionExists(*t) {
		noError = false
		return &noError, "Transition already exists"
	}

	// Add the transition to the map
	a.Transitions[Chemin{start: t.start, letter: t.letter}] = t.end
	return &noError, "Transition added successfully"
}

func (a *Automate) FinishedTransitions() (status *bool, message string) {

	// if there is no transition, return false
	if len(a.Transitions) == 0 {
		noError := false
		return &noError, "Please add at least one transition!!"
	}

	// if there is no transition with a final state, return false
	found := false
	for _, transition := range a.Transitions {
		if a.IsFinalState(transition) {
			found = true
		}
	}

	if !found {
		noError := false
		return &noError, "No transition has final state!!"
	}

	// if there is no transition with an initial state, return false
	result := false
	for chemin := range a.Transitions {
		if chemin.start == a.Q0 && a.IsInitialState(chemin.start) {
			result = true
			break
		}
	}

	if !result {
		noError := false
		return &noError, "No transition with initial state!!"
	}

	noError := true
	return &noError, "Transitions found!"

}
