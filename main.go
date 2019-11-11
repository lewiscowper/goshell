package main

import (
	"bufio"
	"fmt"
	"github.com/lewiscowper/shell/core"
	"github.com/lewiscowper/shell/helpers"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	promptChar := os.Getenv("PROMPTCHAR")
	defaultPromptColor := os.Getenv("PROMPTCOLOR")

	if promptChar == "" {
		promptChar = "❯"
	}

	if defaultPromptColor == "" {
		defaultPromptColor = "cyan"
	}

	exitStatus := 0

	var promptColor string

	var ds []string

	initCurrDir, err := os.Getwd()
	if err != nil {
		exitStatus = 1
		fmt.Fprintln(os.Stderr, err)
	}

	ctx := &helpers.ShellContext{DirStack: helpers.AppendIfMissing(ds, initCurrDir)}

	fmt.Printf("%+v\n", ctx)

	for {
		if exitStatus != 0 {
			promptColor = "red"
		} else {
			promptColor = defaultPromptColor
		}

		displayPrompt := core.BuildPrompt(helpers.PromptParams{Char: promptChar, Color: promptColor})
		fmt.Print(displayPrompt)

		// Read keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input
		err = core.ExecInput(ctx, input)
		if err != nil {
			exitStatus = 1
			fmt.Fprintln(os.Stderr, err)
		} else {
			exitStatus = 0
		}
	}
}
