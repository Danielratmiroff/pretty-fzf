package themes

import (
	"fmt"
	"strings"
)

var themesMap = map[string]func() ColorScheme{
	"catpuccin":            CatpuccinoMocha,
	"catpuccinmocha":       CatpuccinoMocha,
	"catpuccinmacchiato":   CatppuccinMacchiato,
	"catpuccinlatte":       CatppuccinLatte,
	"catpuccinfrappe":      CatppuccinFrappe,
	"gruvbox":              Gruvbox,
	"gruvboxdark":          GruvboxDark,
	"solarizeddark":        SolarizedDark,
	"solarizedlight":       SolarizedLight,
	"tokyonight":           TokyoNight,
	"dracula":              Dracula,
	"github":               GitHub,
	"monokaiextendedlight": MonokaiExtendedLight,
	"sublimesnazzy":        SublimeSnazzy,
}

func SelectTheme(theme string) ColorScheme {
	themeName := strings.ToLower(theme)

	if themeFunc, ok := themesMap[themeName]; ok {
		return themeFunc()
	}

	fmt.Printf("Unknown theme '%s', using default\n", theme)
	return Default()
}

var batcatThemesMap = map[string]string{
	"catpuccin":            "Catppuccin Mocha",
	"catpuccinmocca":       "Catppuccin Mocha",
	"catppuccinmacchiato":  "Catppuccin Macchiato",
	"catppuccinfrappe":     "Catppuccin Frappé",
	"catpuccinlatte":       "Catppuccin Latte",
	"gruvbox":              "Gruvbox",
	"gruvboxdark":          "Gruvbox Dark",
	"solarizeddark":        "Solarized Dark",
	"solarizedlight":       "Solarized Light",
	"tokyonight":           "Tokyo Night",
	"dracula":              "Dracula",
	"github":               "GitHub",
	"monokaiextendedlight": "Monokai Extended Light",
	"sublimesnazzy":        "Sublime Snazzy",
}

func SelectBatcatTheme(theme string) string {
	themeName := strings.ToLower(theme)

	if batcatTheme, ok := batcatThemesMap[themeName]; ok {
		return batcatTheme
	}

	fmt.Printf("Unknown batcat theme '%s', using default\n", theme)
	return ""
}
