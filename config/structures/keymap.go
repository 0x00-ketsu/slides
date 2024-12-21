package structures

// Keymap is the configuration for the key maps
type Keymap struct {
	// Quit is the key to quit the slides.
	Quit []string `mapstructure:"quit"`

	// Help is the key to show the help.
	Help string `mapstructure:"help"`

	// Search is the key to search the slides.
	Search struct {
		// Forward is the key to search forward (downward).
		Forward string `mapstructure:"forward"`

		// PrevMatch is the key to go to the previous match.
		PrevMatch string `mapstructure:"prev-match"`

		// NextMatch is the key to go to the next match.
		NextMatch string `mapstructure:"next-match"`
	}

	// Tagbar is the key to show the tagbar.
	Tagbar struct {
		// Toggle is the key to toggle the tagbar.
		Toggle string `mapstructure:"toggle"`
	}

	// Slide is the key to move the slides.
	Slide struct {
		// First is the key to show the first slide.
		First string `mapstructure:"first"`

		// Last is the key to show the last slide.
		Last string `mapstructure:"last"`

		// Prev is the key to show the previous slide.
		Prev []string `mapstructure:"prev"`

		// Next is the key to show the next slide.
		Next []string `mapstructure:"next"`
	} `mapstructure:"slide"`

	// Move is the key to move the slides.
	Move struct {
		// Top is the key to move to the top.
		Top string `mapstructure:"top"`

		// Bottom is the key to move to the bottom.
		Bottom string `mapstructure:"bottom"`

		// Up is the key to move up.
		Up []string `mapstructure:"up"`

		// Down is the key to move down.
		Down []string `mapstructure:"down"`
	} `mapstructure:"move"`

	// Scroll is the key to scroll the slides.
	Scroll struct {
		// Up is the key to scroll up.
		Up []string `mapstructure:"up"`

		// Down is the key to scroll down.
		Down []string `mapstructure:"down"`
	} `mapstructure:"scroll"`
}
