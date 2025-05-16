package functions

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Clear() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[2J\033[H")
	}
}

func ClearCurrentLine() {
	if runtime.GOOS == "windows" {
		fmt.Print("\r\033[K")
	} else {
		fmt.Print("\033[2K\r")
		fmt.Print("\033[F\033[2K\r")
	}
}

func ClearLine(n ...int) {
	if runtime.GOOS == "windows" {
		if len(n) > 0 {
			for i := 0; i < n[0]; i++ {
				fmt.Print("\r\033[K")
			}
		} else {
			fmt.Print("\r\033[K")
		}
	} else {
		if len(n) > 0 {
			for i := 0; i < n[0]; i++ {
				fmt.Print("\033[F\033[2K\r")
			}
		} else {
			fmt.Print("\033[F\033[2K\r")
		}
	}
}

func Gap() {
	fmt.Println()
	ClearCurrentLine()
	fmt.Println()
	MoveUp()
}

func MoveUp() {
	if runtime.GOOS != "windows" {
		fmt.Print("\033[A")
	}
}
