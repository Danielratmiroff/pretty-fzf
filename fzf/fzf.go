package fzf

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Danielratmiroff/pretty-fzf/themes"
)

type PreviewConfig struct {
	Command string
	Window  string
}

type KeyBinding struct {
	Key    string
	Action string
}

type FZFConfig struct {
	Preview     PreviewConfig
	Colors      themes.ColorScheme
	KeyBindings []KeyBinding
}

func (c FZFConfig) ColorString() string {
	return fmt.Sprintf("bg:%s,fg:%s,header:%s,info:%s,pointer:%s,marker:%s,spinner:%s,prompt:%s,hl:%s",
		c.Colors.Background,
		c.Colors.Foreground,
		c.Colors.Header,
		c.Colors.Info,
		c.Colors.Pointer,
		c.Colors.Marker,
		c.Colors.Spinner,
		c.Colors.Prompt,
		c.Colors.Highlight,
	)
}

func (c FZFConfig) BindString() string {
	var bindings []string
	for _, kb := range c.KeyBindings {
		bindings = append(bindings, fmt.Sprintf("%s:%s", kb.Key, kb.Action))
	}
	return strings.Join(bindings, ",")
}

func (c FZFConfig) ToCommandArgs() []string {
	return []string{
		"--preview", c.Preview.Command,
		"--preview-window", c.Preview.Window,
		"--color", c.ColorString(),
		"--bind", c.BindString(),
	}
}

func NewDefaultConfig() FZFConfig {
	return FZFConfig{
		Preview: PreviewConfig{
			Command: "batcat --style=numbers --color=always --line-range :300 {}",
			Window:  "right:60%",
		},
		Colors: themes.ColorScheme{},
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
