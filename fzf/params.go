package fzf

import (
	"fmt"
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
