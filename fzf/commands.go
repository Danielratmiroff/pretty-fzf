package fzf

import (
	"fmt"
	"strings"

	"github.com/Danielratmiroff/pretty-fzf/themes"
)

type Commands struct {
	PreviewCmd string // Command to run as fzf preview
	RunCmd     string // Command to run to get the output
	OutputCmd  string // Command to run when output is selected
}

var cmdsMap = map[string]func(string) Commands{
	"cd": func(theme string) Commands {
		return changeDirectory()("")
	},
	"findfile": func(theme string) Commands {
		return findFile(theme)("")
	},
}

func SelectCmd(args Params) Commands {
	cmdName := strings.ToLower(args.Cmd)
	theme := themes.SelectBatcatTheme(args.Theme)

	if cmdFunc, ok := cmdsMap[cmdName]; ok {
		return cmdFunc(theme)
	}

	fmt.Printf("Unknown theme '%s', using default\n", args.Cmd)
	return Default()
}

func changeDirectory() func(string) Commands {
	return func(string) Commands {
		return Commands{
			OutputCmd:  "cd",
			PreviewCmd: "'tree -C {} | head -200'",
			RunCmd:     "find . -maxdepth 1 -type d",
		}
	}
}

func findFile(theme string) func(string) Commands {
	return func(string) Commands {
		return Commands{
			OutputCmd:  "cat",
			PreviewCmd: fmt.Sprintf("batcat --theme='%s' --style=numbers --color=always --line-range :300 {}", theme),
			RunCmd:     "fdfind --type f --hidden --follow --exclude .git",
		}
	}
}

func Default() Commands {
	return changeDirectory()("")
}
