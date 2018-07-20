package builtins

import (
	"os/exec"
	"runtime"
)

// Clear tries to clear the user's terminal in an OS agnostic way
func Clear() {
	if runtime.GOOS == "windows" {
		exec.Command("cls")
	} else {
		exec.Command("clear")
	}
}
