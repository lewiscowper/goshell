package helpers

type ShellContext struct {
	DirStack []string
}

// PromptParams are the struct to store the arguments in so I can "provide" default values
type PromptParams struct {
	Char, Color string
}
