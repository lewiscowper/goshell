package core

import (
	"fmt"
	"github.com/mgutz/ansi"

	"github.com/lewiscowper/shell/helpers"
)

// BuildPrompt will build the prompt as needed
func BuildPrompt(p helpers.PromptParams) string {
	if p.Char == "" {
		p.Char = ">"
	}

	if p.Color == "" {
		p.Color = "magenta"
	}

	colorise := ansi.ColorFunc(p.Color)

	return fmt.Sprintf("\n%v ", colorise(p.Char))
}
