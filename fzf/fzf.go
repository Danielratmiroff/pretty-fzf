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

	// Wrap the preview command in single quotes
	previewCmd := fmt.Sprintf("'%s'", config.Preview.Command)
	fzfCmd := fmt.Sprintf("%s | fzf --preview %s", config.Cmd.RunCmd, previewCmd)

	// Build the fzf command
	fzfArgs := append([]string{"-c", fzfCmd}, args...)
	fmt.Printf("Running command: sh %s\n", strings.Join(fzfArgs, " "))

	cmd := exec.Command("sh", fzfArgs...)

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	// Capture the output
	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 130 {
			fmt.Println("fzf was cancelled by the user")
			return nil
		}
		panic(err)
	}

	outputPath := strings.TrimSpace(string(output))

	if outputPath != "" {
		// fmt.Printf("Selected path: %s\n", selectedPath)
		fmt.Printf("%s %s\n", config.Cmd.OutputCmd, outputPath)
	} else {
		panic("No output path selected")
	}

	return nil
}
