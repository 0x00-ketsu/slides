package model

import (
	"bufio"
	"errors"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/0x00-ketsu/slides/internal/styles"
	"github.com/0x00-ketsu/slides/internal/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

const (
	SEP = "\n---\n" // sep of slides
)

// TagBar is the structure for the tagbar.
type TagBar struct {
	// Visible is the visibility of the tagbar.
	Visible bool

	// Width is the width of the tagbar.
	Width   int
}

// Model is the main model for the slides.
type Model struct {
	ready         bool
	viewport      viewport.Model
	help          help.Model
	keys          KeyMap
	quitting      bool
	termRender    glamour.TermRendererOption // glamour style

	// Filename is the markdown file to be loaded
	Filename string

	// Theme is the theme style for rendering markdown
	Theme string

	// Page is the current slide number
	Page int

	// Slide is the current slide content (converted to markdown)
	Slide string

	// Slides are all the slides in the markdown file
	Slides []string

	// TagBar is the tagbar configuration
	TagBar *TagBar

	// Search is the search input
	Search *Search
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
		if m.Search.IsActive {
			switch msg.Type {
			case tea.KeyEnter:
				if m.Search.Query() != "" {
					m.Search.Execute(&m)
				} else {
					m.Search.Done()
				}
				return m, nil
			case tea.KeyCtrlC, tea.KeyEscape:
				m.Search.SetQuery("")
				m.Search.Done()
				return m, nil
			}

			var cmd tea.Cmd
			m.Search.TextInput, cmd = m.Search.TextInput.Update(msg)
			return m, cmd
		}

		switch {
		// Quit
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit

		// Toggle help panel
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil

		// Search
		case key.Matches(msg, DefaultKeyMap.Search):
			m.Search.Begin()
			m.Search.TextInput.Focus()
			return m, nil
		case key.Matches(msg, DefaultKeyMap.NextSearched):
			m.Search.Next(&m)
			return m, nil
		case key.Matches(msg, DefaultKeyMap.PrevSearched):
			m.Search.Previous(&m)
			return m, nil

		// Move
		case key.Matches(msg, DefaultKeyMap.MoveTop):
			m.viewport.GotoTop()
			return m, nil
		case key.Matches(msg, DefaultKeyMap.MoveBottom):
			m.viewport.GotoBottom()
			return m, nil
		case key.Matches(msg, DefaultKeyMap.MoveDown):
			m.viewport.LineDown(1)
			return m, nil
		case key.Matches(msg, DefaultKeyMap.MoveUp):
			m.viewport.LineUp(1)
			return m, nil

		// Scroll
		case key.Matches(msg, DefaultKeyMap.ScrollDown):
			m.viewport.ViewDown()
			return m, nil
		case key.Matches(msg, DefaultKeyMap.ScrollUp):
			m.viewport.ViewUp()
			return m, nil

		// Slide
		case key.Matches(msg, DefaultKeyMap.FirstSlide):
			m.SetPage(0)
			return m, nil
		case key.Matches(msg, DefaultKeyMap.LastSlide):
			m.SetPage(m.PageTotal() - 1)
			return m, nil
		case key.Matches(msg, DefaultKeyMap.PrevSlide):
			prevSlide := m.Page - 1
			if prevSlide < m.PageTotal() {
				m.SetPage(prevSlide)
			}
			return m, nil
		case key.Matches(msg, DefaultKeyMap.NextSlide):
			nextSlide := m.Page + 1
			if nextSlide < m.PageTotal() {
				m.SetPage(nextSlide)
			}
			return m, nil

		// Tagbar
		case key.Matches(msg, DefaultKeyMap.ToggleTagbar):
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
		return errors.New("no slides provided")
	}

	if m.termRender == nil {
		m.termRender = styles.SelectTheme(m.Theme)
	}
	m.help = help.New()
	m.keys = DefaultKeyMap
	slides := m.extractSlides(content)
	slide, _ := m.convertToMarkdown(slides[m.Page])
	m.Slides = slides
	m.Slide = slide
	m.Search = NewSearch()
	return nil
}

// GetAllSlides returns all slides.
func (m *Model) GetAllSlides() []string {
	return m.Slides
}

// GetCurrentSlide returns the current slide.
func (m *Model) GetCurrentPage() int {
	return m.Page
}

// PageTotal returns total count of slides.
func (m *Model) PageTotal() int {
	return len(m.Slides)
}

// Setpage sets the current page.
func (m *Model) SetPage(page int) {
	if page == m.Page || page >= m.PageTotal() || page < 0 {
		return
	}

	slide := m.Slides[page]
	content, _ := m.convertToMarkdown(slide)
	m.Page = page
	m.Slide = slide
	m.viewport.SetContent(content)
}

// Toggle tagbar visibility
func (m *Model) toggleTagbar() {
	m.TagBar.Visible = !m.TagBar.Visible
}

// Convert string to markdown style string.
func (m *Model) convertToMarkdown(slide string) (string, error) {
	r, _ := glamour.NewTermRenderer(m.termRender, glamour.WithWordWrap(m.viewport.Width))
	return r.Render(slide)
}

// Extract slides from markdown content.
func (m *Model) extractSlides(content string) []string {
	content = strings.TrimPrefix(content, strings.TrimPrefix(SEP, "\n"))
	return strings.Split(content, SEP)
}

// Read markdown content from file.
func (m *Model) readFile() (string, error) {
	if !utils.FileExists(m.Filename) {
		return "", errors.New("File is not exist")
	}

	content, err := os.ReadFile(m.Filename)
	return string(content), err
}

// Read markdown content from stdin.
func (m *Model) readStdin() (string, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}

	if stat.Mode()&os.ModeNamedPipe == 0 && stat.Size() == 0 {
		return "", nil
	}

	reader := bufio.NewReader(os.Stdin)
	var b strings.Builder
	for {
		r, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		_, err = b.WriteRune(r)
		if err != nil {
			return "", err
		}
	}
	return b.String(), nil
}

// Extract header text from all slides.
// If not found, return the first unempty line.
func (m *Model) getAllSlidesHeaderText() map[int]string {
	store := make(map[int]string)
	for idx, slide := range m.Slides {
		headerText := m.extractHeaderText(slide)
		if headerText == "" {
			unemptyLine := m.getFirstUnemptyLine(slide)
			if unemptyLine != "" {
				store[idx] = unemptyLine[:20]
			} else {
				store[idx] = ""
			}
		} else {
			store[idx] = headerText
		}
	}
	return store
}

// Get first unempty (stripped) line of plain slide
func (m *Model) getFirstUnemptyLine(plainSlide string) string {
	lines := strings.Split(plainSlide, "\n")
	if len(lines) == 0 {
		return ""
	}

	for _, line := range lines {
		stripLine := strings.TrimSpace(line)
		if len(stripLine) > 0 {
			return stripLine
		}
	}

	return ""
}

// Extract header text from markdown file.
// Return empty string if not found.
func (m *Model) extractHeaderText(plainSlide string) string {
	headerCompile := `^(#+)\s(.*)`
	lines := strings.Split(plainSlide, "\n")
	for _, line := range lines {
		compile, err := regexp.Compile(headerCompile)
		if err != nil {
			return ""
		}

		result := compile.FindStringSubmatch(line)
		matchResultLen := len(result)
		if matchResultLen > 0 {
			return result[len(result)-1]
		}
	}
	return ""
}
