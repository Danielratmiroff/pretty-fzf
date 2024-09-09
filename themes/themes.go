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

func Default() ColorScheme {
	return ColorScheme{
		Background:         "#1E1E2E",
		Foreground:         "#CDD6F4",
		Header:             "#F38BA8",
		Info:               "#CBA6F7",
		Pointer:            "#F5E0DC",
		Marker:             "#F5E0DC",
		Spinner:            "#F5E0DC",
		Prompt:             "#CBA6F7",
		Highlight:          "#F38BA8",
		SelectedForeground: "#CDD6F4",
		SelectedHighlight:  "#F38BA8",
		SelectedBackground: "#313244",
		Border:             "#B4BEFE",
	}
}
