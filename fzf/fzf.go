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
			Command: command.PreviewCmd,
			Window:  "right:60%",
		},
		Colors: theme,
		KeyBindings: []KeyBinding{
			{Key: "ctrl-d", Action: "preview-down"},
			{Key: "ctrl-u", Action: "preview-up"},
			{Key: "ctrl-f", Action: "preview-page-down"},
			{Key: "ctrl-b", Action: "preview-page-up"},
		},
		Cmd: command,
	}
}

func RunCommand(config FZFConfig) error {
	args := config.ToCommandArgs()

	// Build main(run) command
	fullCmd := fmt.Sprintf("%s | fzf %s", config.Cmd.RunCmd, strings.Join(args, " "))
	fmt.Println("Executing command:", fullCmd)

	fzfCmd := exec.Command(fullCmd)

	fzfCmd.Stdin = os.Stdin
	fzfCmd.Stderr = os.Stderr
	// Capture the output of fzf
	output, err := fzfCmd.Output()
	if err != nil {
		return err
	}

	selectedPath := strings.TrimSpace(string(output))

	// Execute output command
	cdCmd := exec.Command(config.Cmd.OutputCmd, selectedPath)

	fmt.Printf("Changing directory to: %s\n", selectedPath)

	return cdCmd.Run()
}
