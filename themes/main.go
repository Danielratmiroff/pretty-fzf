package themes

import "fmt"

var themesMap = map[string]func() ColorScheme{
	"catpuccin":          CatpuccinoMocha,
	"catpuccinMocha":     CatpuccinoMocha,
	"catpuccinMacchiato": CatppuccinMacchiato,
	"catpuccinLatte":     CatppuccinLatte,
	"catpuccinFrappe":    CatppuccinFrappe,
	"gruvbox":            Gruvbox,
}

func SelectTheme(theme string) ColorScheme {
	if themeFunc, ok := themesMap[theme]; ok {
		return themeFunc()
	}
	fmt.Printf("Unknown theme '%s', using default\n", theme)
	return Default()
}
