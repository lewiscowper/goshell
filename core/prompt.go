package core

import (
	"fmt"
	"github.com/mgutz/ansi"
)

// PromptParams are the struct to store the arguments in so I can "provide" default values
type PromptParams struct {
	Char, Color string
}

// BuildPrompt will build the prompt as needed
func BuildPrompt(p PromptParams) string {
	if p.Char == "" {
		p.Char = ">"
	}

	if p.Color == "" {
		p.Color = "magenta"
	}

	colorise := ansi.ColorFunc(p.Color)

	return fmt.Sprintf("\n%v ", colorise(p.Char))
}
