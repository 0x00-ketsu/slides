package model

import (
	"github.com/0x00-ketsu/slides/config"
	"github.com/charmbracelet/bubbles/key"
)

var k = config.LoadKeymaps()

// Keymap represents the keymap configuration.
type KeyMap struct {
	Help key.Binding
	Quit key.Binding

	FirstSlide key.Binding
	LastSlide  key.Binding
	PrevSlide  key.Binding
	NextSlide  key.Binding

	MoveTop    key.Binding
	MoveBottom key.Binding
	MoveUp     key.Binding
	MoveDown   key.Binding

	ScrollDown key.Binding
	ScrollUp   key.Binding

	Search       key.Binding
	NextSearched key.Binding
	PrevSearched key.Binding

	ToggleTagbar key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view.
// It's part of the key.Map interface.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help}
}

// FullHelp returns keybindings for the expanded help view.
// It's part of the key.Map interface.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit},
		{k.FirstSlide, k.LastSlide, k.NextSlide, k.PrevSlide},
		{k.MoveTop, k.MoveBottom, k.MoveUp, k.MoveDown},
		{k.ScrollDown, k.ScrollUp},
		{k.Search, k.NextSearched, k.PrevSearched},
		{k.ToggleTagbar},
	}
}

var DefaultKeyMap = KeyMap{
	Help: key.NewBinding(
		key.WithKeys(k.Help),
		key.WithHelp(k.Help, "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys(k.Quit...),
		key.WithHelp(k.Quit[0], "quit"),
	),

	// Slide
	FirstSlide: key.NewBinding(
		key.WithKeys(k.Slide.First),
		key.WithHelp(k.Slide.First, "first slide"),
	),
	LastSlide: key.NewBinding(
		key.WithKeys(k.Slide.Last),
		key.WithHelp(k.Slide.Last, "last slide"),
	),
	NextSlide: key.NewBinding(
		key.WithKeys(k.Slide.Next...),
		key.WithHelp(k.Slide.Next[0], "next slide"),
	),
	PrevSlide: key.NewBinding(
		key.WithKeys(k.Slide.Prev...),
		key.WithHelp(k.Slide.Prev[0], "previous slide"),
	),

	// Move
	MoveTop: key.NewBinding(
		key.WithKeys(k.Move.Top),
		key.WithHelp(k.Move.Top, "move to the top"),
	),
	MoveBottom: key.NewBinding(
		key.WithKeys(k.Move.Bottom),
		key.WithHelp(k.Move.Bottom, "move to the bottom"),
	),
	MoveUp: key.NewBinding(
		key.WithKeys(k.Move.Up...),
		key.WithHelp(k.Move.Up[0], "move up"),
	),
	MoveDown: key.NewBinding(
		key.WithKeys(k.Move.Down...),
		key.WithHelp(k.Move.Down[0], "move down"),
	),

	// scroll
	ScrollDown: key.NewBinding(
		key.WithKeys(k.Scroll.Down...),
		key.WithHelp(k.Scroll.Down[0], "scroll down"),
	),
	ScrollUp: key.NewBinding(
		key.WithKeys(k.Scroll.Up...),
		key.WithHelp(k.Scroll.Up[0], "scroll up"),
	),

	// Search
	Search: key.NewBinding(
		key.WithKeys(k.Search.Forward),
		key.WithHelp(k.Search.Forward, "search forward"),
	),
	PrevSearched: key.NewBinding(
		key.WithKeys(k.Search.PrevMatch),
		key.WithHelp(k.Search.PrevMatch, "go to the previous match in the slides"),
	),
	NextSearched: key.NewBinding(
		key.WithKeys(k.Search.NextMatch),
		key.WithHelp(k.Search.NextMatch, "go to the next match in the slides"),
	),

	// Tagbar
	ToggleTagbar: key.NewBinding(
		key.WithKeys(k.Tagbar.Toggle),
		key.WithHelp(k.Tagbar.Toggle, "toggle tagbar"),
	),
}
