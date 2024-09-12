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
	Cmd   string
}

func NewDefaultConfig(params Params) FZFConfig {
	theme := themes.SelectTheme(params.Theme)
	command := SelectCmd(params)

	return FZFConfig{
		Preview: PreviewConfig{
			Command: command.preview,
			Window:  "right:60%",
		},
		Colors: theme,
		KeyBindings: []KeyBinding{
			{Key: "ctrl-d", Action: "preview-down"},
			{Key: "ctrl-u", Action: "preview-up"},
			{Key: "ctrl-f", Action: "preview-page-down"},
			{Key: "ctrl-b", Action: "preview-page-up"},
		},
		Cmd: params.Cmd,
	}
}

func RunCommand(config FZFConfig) error {
	args := config.ToCommandArgs()

	fzfCmd := exec.Command("fish", "-c", "find . -maxdepth 1 -type d | fzf "+strings.Join(args, " "))

	fzfCmd.Stdin = os.Stdin
	fzfCmd.Stderr = os.Stderr

	// Capture the output of fzf
	output, err := fzfCmd.Output()
	if err != nil {
		return err
	}

	fullCommand := "fzf " + strings.Join(args, " ")
	fmt.Println("Executing command:", fullCommand)

	selectedPath := strings.TrimSpace(string(output))

	// Execute the cd command
	cdCmd := exec.Command(config.Cmd, selectedPath)

	fmt.Printf("Changing directory to: %s\n", selectedPath)

	return cdCmd.Run()
}
