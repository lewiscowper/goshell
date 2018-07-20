package core

import (
	"github.com/lewiscowper/shell/builtins"
	"os"
	"os/exec"
	"strings"
)

// ExecInput takes an input string and executes it
func ExecInput(input string) error {
	// Remove the newline character
	input = strings.TrimSuffix(input, "\n")

	// Split the input to separate the command and the arguments
	args := strings.Split(input, " ")

	// Check for built-in commands
	switch args[0] {
	case "cd":
		err := builtins.Cd(args)
		if err != nil {
			return err
		}

		return nil
	case "clear":
		builtins.Clear()
	case "exit":
		os.Exit(0)
	}

	// Prepare the command to execute
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and save the output
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
