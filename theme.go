package c4

// Theme holds the top-level theming information for the diagram.
type Theme struct {
	// Default styles for all system elements.
	System Palette

	// Default styles for all container elements.
	Container Palette

	// Default styles for all component elements.
	Component Palette

	// Default styles for all person elements.
	Person Palette
}

// Palette holds individual theming parameters.
type Palette struct {
	// The background color of the element.
	BackgroundColor string

	// The font color for text within the element.
	FontColor string
}

// DefaultTheme returns the styles used for diagrams without an explicit theme.
func DefaultTheme() Theme {
	return Theme{
		System: Palette{
			BackgroundColor: "#4E668A",
			FontColor:       "#F5F5F5",
		},
		Container: Palette{
			BackgroundColor: "#6C8EBF",
			FontColor:       "#262626",
		},
		Component: Palette{
			BackgroundColor: "#94B3E0",
			FontColor:       "#262626",
		},
		Person: Palette{
			BackgroundColor: "#455A7A",
			FontColor:       "#ffffff",
		},
	}
}
