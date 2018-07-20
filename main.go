package main

import (
	"bufio"
	"fmt"
	"github.com/mgutz/ansi"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"github.com/lewiscowper/shell/builtins"
)

func clear() {
	if runtime.GOOS == "windows" {
		exec.Command("cls")
	} else {
		exec.Command("clear")
	}
}

func execInput(input string) error {
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
		clear()
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

type promptParams struct {
	char, color string
}

func buildPrompt(p promptParams) string {
	if p.char == "" {
		p.char = ">"
	}

	if p.color == "" {
		p.color = "magenta"
	}

	colourise := ansi.ColorFunc(p.color)

	return fmt.Sprintf("\n%v ", colourise(p.char))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	promptChar := os.Getenv("PROMPTCHAR")
	promptColor := os.Getenv("PROMPTCOLOR")

	exitStatus := 0

	for {
		if exitStatus != 0 {
			promptColor = "red"
		} else {
			promptColor = os.Getenv("PROMPTCOLOR")
		}

		fmt.Print(buildPrompt(promptParams{char: promptChar, color: promptColor}))

		// Read keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		// Handle the execution of the input
		err = execInput(input)
		if err != nil {
			exitStatus = 1
			fmt.Println(err)
		} else {
			exitStatus = 0
		}
	}
}
