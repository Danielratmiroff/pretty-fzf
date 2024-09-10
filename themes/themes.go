package themes

type ColorScheme struct {
	Background         string
	Foreground         string
	Header             string
	Info               string
	Pointer            string
	Marker             string
	Spinner            string
	Prompt             string
	Highlight          string
	SelectedForeground string
	SelectedHighlight  string
	SelectedBackground string
	Border             string
}

func Default() ColorScheme {
	return Dracula()
}

func Gruvbox() ColorScheme {
	return ColorScheme{
		Background: "#1E1E2E",
		Foreground: "#CDD6F4",
		Header:     "#F38BA8",
		Info:       "#CBA6F7",
		Pointer:    "#F5E0DC",
		Marker:     "#F5E0DC",
		Spinner:    "#F5E0DC",
		Prompt:     "#CBA6F7",
		Highlight:  "#F38BA8",
	}
}

func GruvboxDark() ColorScheme {
	return ColorScheme{
		Background:         "#282828", // Dark background
		Foreground:         "#ebdbb2", // Light foreground
		Header:             "#fb4934", // Red
		Info:               "#83a598", // Green
		Pointer:            "#d79921", // Yellow
		Marker:             "#fe8019", // Orange
		Spinner:            "#b16286", // Purple
		Prompt:             "#8ec07c", // Light Green
		Highlight:          "#d65d0e", // Bright Orange
		SelectedForeground: "#fbf1c7", // Light background for selected text
		SelectedHighlight:  "#b8bb26", // Bright Yellow
		SelectedBackground: "#3c3836", // Darker background for selection
		Border:             "#7c6f64", // Soft Brown
	}
}

func CatpuccinoMocha() ColorScheme {
	return ColorScheme{
		Background:         "#1E1E2E", // Base
		Foreground:         "#CDD6F4", // Text
		Header:             "#F5C2E7", // Pink
		Info:               "#89B4FA", // Blue
		Pointer:            "#F5E0DC", // Rosewater
		Marker:             "#FAB387", // Peach
		Spinner:            "#94E2D5", // Teal
		Prompt:             "#CBA6F7", // Mauve
		Highlight:          "#F38BA8", // Red
		SelectedForeground: "#1E1E2E", // Base (inverted for selection)
		SelectedHighlight:  "#F9E2AF", // Yellow
		SelectedBackground: "#585B70", // Surface2
		Border:             "#B4BEFE", // Lavender
	}
}

func CatppuccinMacchiato() ColorScheme {
	return ColorScheme{
		Background:         "#24273A", // Base
		Foreground:         "#CAD3F5", // Text
		Header:             "#F4DBD6", // Rosewater
		Info:               "#8AADF4", // Blue
		Pointer:            "#F0C6C6", // Flamingo
		Marker:             "#F5A97F", // Peach
		Spinner:            "#8BD5CA", // Teal
		Prompt:             "#C6A0F6", // Mauve
		Highlight:          "#ED8796", // Red
		SelectedForeground: "#24273A", // Base (inverted for selection)
		SelectedHighlight:  "#EED49F", // Yellow
		SelectedBackground: "#5B6078", // Surface2
		Border:             "#B7BDF8", // Lavender
	}
}

func CatppuccinFrappe() ColorScheme {
	return ColorScheme{
		Background:         "#303446", // Base
		Foreground:         "#C6D0F5", // Text
		Header:             "#F2D5CF", // Rosewater
		Info:               "#8CAAEE", // Blue
		Pointer:            "#EEBEBE", // Flamingo
		Marker:             "#EF9F76", // Peach
		Spinner:            "#81C8BE", // Teal
		Prompt:             "#CA9EE6", // Mauve
		Highlight:          "#E78284", // Red
		SelectedForeground: "#303446", // Base (inverted for selection)
		SelectedHighlight:  "#E5C890", // Yellow
		SelectedBackground: "#626880", // Surface2
		Border:             "#BABBF1", // Lavender
	}
}

func CatppuccinLatte() ColorScheme {
	return ColorScheme{
		Background:         "#EFF1F5", // Base
		Foreground:         "#4C4F69", // Text
		Header:             "#DC8A78", // Maroon
		Info:               "#1E66F5", // Blue
		Pointer:            "#DD7878", // Red
		Marker:             "#FE640B", // Peach
		Spinner:            "#179299", // Teal
		Prompt:             "#8839EF", // Mauve
		Highlight:          "#D20F39", // Red
		SelectedForeground: "#EFF1F5", // Base (inverted for selection)
		SelectedHighlight:  "#DF8E1D", // Yellow
		SelectedBackground: "#ACB0BE", // Surface2
		Border:             "#7287FD", // Lavender
	}
}

func SolarizedDark() ColorScheme {
	return ColorScheme{
		Background:         "#002b36", // Base background
		Foreground:         "#839496", // Base text
		Header:             "#dc322f", // Red
		Info:               "#268bd2", // Blue
		Pointer:            "#b58900", // Yellow
		Marker:             "#cb4b16", // Orange
		Spinner:            "#d33682", // Magenta
		Prompt:             "#2aa198", // Cyan
		Highlight:          "#859900", // Green
		SelectedForeground: "#93a1a1", // Light foreground for selected text
		SelectedHighlight:  "#b58900", // Yellow for selection highlight
		SelectedBackground: "#073642", // Darker background for selection
		Border:             "#586e75", // Base border color
	}
}

func TokyoNight() ColorScheme {
	return ColorScheme{
		Background:         "#1A1B26", // Dark background
		Foreground:         "#B0B3CC", // Light foreground
		Header:             "#FF6C6B", // Red
		Info:               "#7AA2F7", // Blue
		Pointer:            "#F7768E", // Pink
		Marker:             "#E0AF68", // Orange
		Spinner:            "#9ECE6A", // Green
		Prompt:             "#C0CAF5", // Light Purple
		Highlight:          "#BB9AF7", // Purple
		SelectedForeground: "#1F1F28", // Darker background for selected text
		SelectedHighlight:  "#F7768E", // Pink for selection highlight
		SelectedBackground: "#2A2A3D", // Slightly lighter background for selection
		Border:             "#3B3B58", // Soft border color
	}
}

func Dracula() ColorScheme {
	return ColorScheme{
		Background:         "#282a36", // Dark background
		Foreground:         "#f8f8f2", // Light foreground
		Header:             "#ff79c6", // Pink
		Info:               "#8be9fd", // Cyan
		Pointer:            "#50fa7b", // Green
		Marker:             "#ffb86c", // Orange
		Spinner:            "#bd93f9", // Purple
		Prompt:             "#ff5555", // Red
		Highlight:          "#ff79c6", // Pink
		SelectedForeground: "#f8f8f2", // Light foreground for selected text
		SelectedHighlight:  "#ffb86c", // Orange for selection highlight
		SelectedBackground: "#44475a", // Darker background for selection
		Border:             "#6272a4", // Soft blue-grey
	}
}

func GitHub() ColorScheme {
	return ColorScheme{
		Background:         "#ffffff", // White background
		Foreground:         "#24292e", // Dark text
		Header:             "#0366d6", // Blue for headers
		Info:               "#6f42c1", // Purple for info
		Pointer:            "#d73a49", // Red for pointers
		Marker:             "#f6c94c", // Yellow for markers
		Spinner:            "#28a745", // Green for spinners
		Prompt:             "#24292e", // Dark text for prompts
		Highlight:          "#f1e05a", // Yellow for highlights
		SelectedForeground: "#ffffff", // White for selected text
		SelectedHighlight:  "#0366d6", // Blue for selection highlight
		SelectedBackground: "#f6f8fa", // Light grey for selection background
		Border:             "#e1e4e8", // Light grey for borders
	}
}

func SolarizedLight() ColorScheme {
	return ColorScheme{
		Background:         "#fdf6e3", // Base background
		Foreground:         "#657b83", // Base text
		Header:             "#dc322f", // Red
		Info:               "#268bd2", // Blue
		Pointer:            "#b58900", // Yellow
		Marker:             "#cb4b16", // Orange
		Spinner:            "#d33682", // Magenta
		Prompt:             "#2aa198", // Cyan
		Highlight:          "#859900", // Green
		SelectedForeground: "#586e75", // Darker foreground for selected text
		SelectedHighlight:  "#b58900", // Yellow for selection highlight
		SelectedBackground: "#eee8d5", // Light background for selection
		Border:             "#93a1a1", // Base border color
	}
}

func MonokaiExtendedLight() ColorScheme {
	return ColorScheme{
		Background:         "#F8F8F2", // Light background
		Foreground:         "#272822", // Dark text
		Header:             "#F92672", // Pink for headers
		Info:               "#66D9EF", // Cyan for info
		Pointer:            "#F8F8F2", // White for pointers
		Marker:             "#FD971F", // Orange for markers
		Spinner:            "#A6E22E", // Green for spinners
		Prompt:             "#F92672", // Pink for prompts
		Highlight:          "#E6DB74", // Yellow for highlights
		SelectedForeground: "#272822", // Dark text for selected text
		SelectedHighlight:  "#F92672", // Pink for selection highlight
		SelectedBackground: "#E6DB74", // Light yellow for selection background
		Border:             "#C5C8C6", // Soft grey for borders
	}
}

func SublimeSnazzy() ColorScheme {
	return ColorScheme{
		Background:         "#272822", // Dark background
		Foreground:         "#F8F8F2", // Light foreground
		Header:             "#F92672", // Pink for headers
		Info:               "#66D9EF", // Cyan for info
		Pointer:            "#F8F8F2", // White for pointers
		Marker:             "#FD971F", // Orange for markers
		Spinner:            "#A6E22E", // Green for spinners
		Prompt:             "#F92672", // Pink for prompts
		Highlight:          "#E6DB74", // Yellow for highlights
		SelectedForeground: "#F8F8F2", // Light foreground for selected text
		SelectedHighlight:  "#F92672", // Pink for selection highlight
		SelectedBackground: "#44475A", // Darker background for selection
		Border:             "#C5C8C6", // Soft grey for borders
	}
}
