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

func Catpuccino() ColorScheme {
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
