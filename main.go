package main

import (
	"bufio"
	"fmt"
	"github.com/lewiscowper/shell/core"
	"os"
	"strings"
)

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

		displayPrompt := core.BuildPrompt(core.PromptParams{Char: promptChar, Color: promptColor})
		fmt.Print(displayPrompt)

		// Read keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		// Remove the newline character
		input = strings.TrimSuffix(input, "\n")

		// Handle the execution of the input
		err = core.ExecInput(input)
		if err != nil {
			exitStatus = 1
			fmt.Println(err)
		} else {
			exitStatus = 0
		}
	}
}
