package helper

import "github.com/charmbracelet/bubbles/key"

// KeyMap defines a set of keybindings. To work for help it must satisfy key.Map.
// It could also very easily be a map[string]key.Binding.
type KeyMap struct {
	NextSlide     key.Binding
	PreviousSlide key.Binding
	FirstSlide    key.Binding
	LastSlide     key.Binding

	ScrollDown key.Binding
	ScrollUp   key.Binding

	Search           key.Binding
	NextSearched     key.Binding
	PreviousSearched key.Binding

	Tagbar key.Binding

	Help key.Binding
	Quit key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view.
// It's part of the key.Map interface.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.NextSlide, k.PreviousSlide, k.FirstSlide, k.LastSlide},
		{k.ScrollDown, k.ScrollUp},
		{k.Search, k.NextSearched, k.PreviousSearched},
		{k.Tagbar},
		{k.Help, k.Quit},
	}
}

var Keys = KeyMap{
	// slide
	NextSlide: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "next slide"),
	),
	PreviousSlide: key.NewBinding(
		key.WithKeys("h"),
		key.WithHelp("h", "previous slide"),
	),
	FirstSlide: key.NewBinding(
		key.WithKeys("g"),
		key.WithHelp("g", "first slide"),
	),
	LastSlide: key.NewBinding(
		key.WithKeys("G"),
		key.WithHelp("G", "last slide"),
	),

	// scroll
	ScrollDown: key.NewBinding(
		key.WithKeys("j"),
		key.WithHelp("j", "scroll down slide"),
	),
	ScrollUp: key.NewBinding(
		key.WithKeys("k"),
		key.WithHelp("k", "scroll up slide"),
	),

	// search
	Search: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "search word"),
	),
	NextSearched: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "next search mathed slide"),
	),
	PreviousSearched: key.NewBinding(
		key.WithKeys("N"),
		key.WithHelp("N", "previous search mathed slide"),
	),

	// tagbar
	Tagbar: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "toggle tagbar"),
	),

	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
