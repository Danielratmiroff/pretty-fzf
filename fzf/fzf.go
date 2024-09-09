package fzf

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Danielratmiroff/pretty-fzf/themes"
)

type Params struct {
	Theme string
}

func NewDefaultConfig(params Params) FZFConfig {
	theme := themes.SelectTheme(params.Theme)
	return FZFConfig{
		Preview: PreviewConfig{
			Command: "batcat --theme='Catppuccin Mocha' --style=numbers --color=always --line-range :300 {}",
			Window:  "right:60%",
		},
		Colors: theme,
		KeyBindings: []KeyBinding{
			{Key: "ctrl-d", Action: "preview-down"},
			{Key: "ctrl-u", Action: "preview-up"},
			{Key: "ctrl-f", Action: "preview-page-down"},
			{Key: "ctrl-b", Action: "preview-page-up"},
		},
	}
}

func RunCommand(config FZFConfig) error {
	args := config.ToCommandArgs()

	// Prepare the command
	cmd := exec.Command("fzf", args...)

	// Set up the command's standard input, output, and error
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Create a single string of all arguments for debugging
	fullCommand := "fzf " + strings.Join(args, " ")
	fmt.Println("Executing command:", fullCommand)

	return cmd.Run()
}
