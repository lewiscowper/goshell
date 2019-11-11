package core

import (
	"errors"
	"fmt"
	"github.com/lewiscowper/shell/builtins"
	"github.com/lewiscowper/shell/helpers"
	"os"
	"os/exec"
	"strings"
)

// ExecInput takes an input string and executes it
func ExecInput(ctx *helpers.ShellContext, input string) error {
	// Remove the newline character
	input = strings.TrimSuffix(input, "\n")

	// Split the input to separate the command and the arguments
	args := strings.Split(input, " ")

	// Check for built-in commands
	switch args[0] {
	case "cd":
		err := builtins.Cd(ctx, args[1:])
		if err != nil {
			return err
		}

		return nil
	case "clear":
		builtins.Clear()
	case "exit":
		os.Exit(0)
	}

	pathToBin, err := exec.LookPath(args[0])
	if err != nil {
		return errors.New(fmt.Sprintf("%s not found in your PATH", args[0]))
	}

	// Prepare the command to execute
	cmd := exec.Command(pathToBin, args[1:]...)

	// Set the correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Set all environment variables from the shell
	cmd.Env = os.Environ()

	// Execute the command and save the output
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
