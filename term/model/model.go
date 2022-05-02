package model

import (
	"errors"
	"strconv"

	"github.com/0x00-ketsu/smooth/term/helper"
	"github.com/0x00-ketsu/smooth/term/search"
	"github.com/0x00-ketsu/smooth/term/styles"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

const (
	SEP = "\n---\n" // sep of slides
)

type Model struct {
	ready        bool           // is file loaded
	activeTagbar bool           // is show tagbar
	viewport     viewport.Model // bubbles viewport model
	help         help.Model
	keys         helper.KeyMap
	quitting     bool

	Filename     string                     // markdown file name
	Theme        string                     // markdown theme style
	Page         int                        // current visit slide number, from 0
	Slide        string                     // current visit slide content (converted to markdown)
	Slides       []string                   // all original slides
	Search       search.Search              // search input
	termRenderer glamour.TermRendererOption // glamour style
}

func (m Model) Init() tea.Cmd {
	// TODO: refresh slides when file is updated
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		keyPress := msg.String()
		if m.Search.Active {
			switch msg.Type {
			case tea.KeyEnter:
				// execute current buffer
				if m.Search.Query() != "" {
					m.Search.Execute(&m)
				} else {
					m.Search.Done()
				}
				// cancel search
				return m, nil
			case tea.KeyCtrlC, tea.KeyEscape:
				// quit command mode
				m.Search.SetQuery("")
				m.Search.Done()
				return m, nil
			}

			var cmd tea.Cmd
			m.Search.SearchTextInput, cmd = m.Search.SearchTextInput.Update(msg)
			return m, cmd
		}

		switch {
		// quit app
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		// help pane
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		}

		switch keyPress {
		// forward Search
		case "/":
			m.Search.Begin()
			m.Search.SearchTextInput.Focus()
			return m, nil

		// goto next searched slide if exist
		case "n":
			m.Search.Next(&m)
			return m, nil

		// goto previous searched slide if exist
		case "N":
			m.Search.Previous(&m)
			return m, nil

		// scroll down when slide is over than terminal height
		case "j":
			m.viewport.ViewDown()
			return m, nil

		// scroll up when slide is over than terminal height
		case "k":
			m.viewport.ViewUp()
			return m, nil

		// previous slide
		case "h":
			previousPage := m.Page - 1
			if previousPage < m.PageSize() {
				m.SetPage(previousPage)
			}
			return m, nil

		// next slide
		case "l":
			nextPage := m.Page + 1
			if nextPage < m.PageSize() {
				m.SetPage(nextPage)
			}
			return m, nil

		// goto specific Slide
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			page, err := strconv.Atoi(keyPress)
			if err == nil {
				m.SetPage(page - 1)
			}
			return m, nil

		// goto first slide
		case "g":
			m.SetPage(0)
			return m, nil

		// goto last slide
		case "G":
			m.SetPage(m.PageSize() - 1)
			return m, nil

		// toggle tagbar
		case "t":
			m.toggleTagbar()
			return m, nil
		}

	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-1)
			m.viewport.SetContent(m.Slide)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - 1
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}

	return m.Render()
}

// Initial does some initial actions
func (m *Model) Initial() error {
	var (
		content string
		err     error
	)

	if m.Filename != "" {
		content, err = m.readFile()
	} else {
		content, err = m.readStdin()
	}

	if err != nil {
		return err
	}

	if content == "" {
		return errors.New("no slides provided\n")
	}

	if m.termRenderer == nil {
		m.termRenderer = styles.SelectTheme(m.Theme)
	}

	m.help = help.New()
	m.keys = helper.Keys

	slides := m.extracSlides(content)
	slide, _ := m.convertToMarkdown(slides[m.Page])

	m.Slides = slides
	m.Slide = slide
	m.Search = search.NewSearch()

	return nil
}

// GetAllSlides returns all original slides
func (m *Model) GetAllSlides() []string {
	return m.Slides
}

func (m *Model) CurrentPage() int {
	return m.Page
}

// PageSize returns total slides count
func (m *Model) PageSize() int {
	return len(m.Slides)
}

func (m *Model) SetPage(page int) {
	if page == m.Page || page >= m.PageSize() || page < 0 {
		return
	}

	slide := m.Slides[page]
	content, _ := m.convertToMarkdown(slide)

	m.Page = page
	m.Slide = slide
	m.viewport.SetContent(content)
}

// toggle tagbar (show/hide tagbar)
func (m *Model) toggleTagbar() {
	m.activeTagbar = !m.activeTagbar
}
