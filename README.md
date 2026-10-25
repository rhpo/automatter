# Automatter

**Automatter** is an interactive command-line tool for defining and simulating **Deterministic Finite Automata (DFA)**. It lets you build an automaton from scratch by specifying its formal components through a guided interface, then instantly test whether arbitrary words are accepted by the language it recognizes.

---

## Theoretical Background

A **Deterministic Finite Automaton** is a 5-tuple:

$$M = (Q, \Sigma, \delta, q_0, F)$$

where:

| Component | Description |
|-----------|-------------|
| $Q$ | A finite, non-empty set of **states** |
| $\Sigma$ | A finite **alphabet** (set of input symbols) |
| $\delta: Q \times \Sigma \to Q$ | A **transition function** mapping a state and a symbol to the next state |
| $q_0 \in Q$ | The **initial state** |
| $F \subseteq Q$ | A set of **accepting (final) states** |

A word $w = a_1 a_2 \cdots a_n \in \Sigma^*$ is **accepted** by $M$ if and only if the sequence of transitions starting from $q_0$ leads to a state in $F$ after consuming all symbols. Formally, $M$ accepts $w$ iff $\hat\delta(q_0, w) \in F$, where $\hat\delta$ is the extended transition function.

Automatter implements this model faithfully: the transition relation is stored as a map $\delta: (Q \times \Sigma) \to Q$, ensuring at most one target state per (state, symbol) pair, which is the defining property of determinism.

---

## Features

- Interactive, step-by-step construction of a DFA through the terminal
- Full validation at each input step (duplicate states, unknown symbols, missing transitions, etc.)
- Colored tabular display of the transition table $\delta$
- Real-time word acceptance testing in a read-eval loop
- Cross-platform terminal control (Windows `cmd` and POSIX ANSI escape codes)

---

## Project Structure

```
automatter/
├── main.go              # Entry point: orchestrates construction, display, and evaluation loop
├── type.go              # Core data model: Automate, Transition, Chemin types and all methods
├── read.go              # Interactive input layer: reads alphabet, states, and transitions from stdin
├── eval.go              # Evaluation engine: word acceptance via transition traversal
├── display.go           # Rendering layer: prints the automaton summary and transition table
└── functions/
    └── functions.go     # Terminal utility functions: clear screen, clear line, cursor movement
```

### Component Responsibilities

**`type.go` — Data Model**

Defines the `Automate` struct and all mutation/query methods. The transition function $\delta$ is represented as:

```go
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
```

A `Chemin` (French for *path*) encodes a (state, symbol) pair, making `Transitions` a direct Go-level implementation of $\delta$. All state and symbol additions are guarded by membership checks to keep the automaton consistent at all times.

**`read.go` — Input Layer**

Guides the user through defining each component of the 5-tuple in order: $\Sigma$, $Q$, $q_0$, $F$, then $\delta$. Each field is read in a validation loop that rejects invalid or duplicate entries and re-prompts on error, so the resulting automaton is always well-formed before evaluation begins.

**`eval.go` — Acceptance Engine**

Implements word acceptance by simulating $\hat\delta$ iteratively. For each symbol in the input word, the engine looks up the unique successor state via the `Transitions` map. If any symbol has no defined transition from the current state, the word is immediately rejected. Acceptance is confirmed if and only if the final state reached belongs to $F$.

**`display.go` — Presentation Layer**

Renders the automaton as a formatted transition table, with states as rows and alphabet symbols as columns. Uses ANSI color codes (via `fatih/color`) to visually distinguish borders, headers, and cell values.

**`functions/functions.go` — Terminal Utilities**

Provides cross-platform helpers for clearing the screen and erasing lines, abstracting over Windows `cls` and POSIX escape sequences so the interactive UI stays clean across environments.

---

## Getting Started

### Prerequisites

- Go 1.22 or later

### Build and Run

```bash
git clone https://github.com/yourusername/automatter.git
cd automatter
go run .
```

Or build a binary:

```bash
go build -o automatter .
./automatter
```

### Usage

The program walks you through defining the automaton interactively:

1. **Alphabet** $\Sigma$: enter symbols one by one, press Enter on an empty line to finish.
2. **States** $Q$: enter state names one by one, press Enter to finish.
3. **Initial state** $q_0$: enter one of the declared states.
4. **Final states** $F$: enter one or more states from $Q$, press Enter to finish.
5. **Transitions** $\delta$: for each transition, enter the source state, the symbol, and the target state. Press Enter on an empty source to finish.

Once constructed, the automaton is displayed as a transition table. The program then enters an interactive loop where you can type any word and receive an immediate accept/reject verdict.

---

## Dependencies

| Package | Purpose |
|---------|---------|
| [`github.com/fatih/color`](https://github.com/fatih/color) | Cross-platform ANSI color output |
| `golang.org/x/sys` | Low-level OS primitives (transitive dependency) |

---

## License

This project is open source. See [LICENSE](LICENSE) for details.

