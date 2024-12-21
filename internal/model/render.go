package model

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/0x00-ketsu/slides/internal/styles"
	"github.com/charmbracelet/lipgloss"
)

// Render returns the rendered view of the model.
func (m *Model) Render() string {
	scrollPercent := m.viewport.ScrollPercent() * 100
	if math.IsNaN(scrollPercent) {
		scrollPercent = 0
	}
	status := fmt.Sprintf("scroll %3.f%%", scrollPercent)

	// Main view
	var main string
	if m.TagBar.Visible {
		main = lipgloss.JoinHorizontal(
			lipgloss.Top,
			styles.Tagbar.Render(m.renderTagbar()),
			m.viewport.View(),
		)
	} else {
		main = lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.viewport.View(),
		)
	}

	// Bottom status bar
	var bottom string
	helpView := m.help.View(m.keys)
	if m.Search.IsActive {
		bottom = m.Search.TextInput.View()
	} else {
		bottom = m.joinHorizontal(
			m.viewport.Width,
			helpView,
			styles.Status.Render(status),
			styles.Page.Render(fmt.Sprintf("slide %d / %d", m.GetCurrentPage()+1, m.PageTotal())),
		)
	}
	return m.joinVertical(m.viewport.Height+2, main, bottom)
}

func (m *Model) renderTagbar() string {
	slidesHeaderText := m.getAllSlidesHeaderText()
	keys := make([]int, 0, len(slidesHeaderText))
	for k := range slidesHeaderText {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var (
		lines      []string
		line       string
		maxLineLen int
	)
	for _, val := range keys {
		if val == m.GetCurrentPage() {
			line = fmt.Sprintf("󰓾 %d: %s", val+1, slidesHeaderText[val])
			maxLineLen = max(maxLineLen, len(line))
			line = styles.ActiveTagbar.Render(line)
		} else {
			line = fmt.Sprintf("%d: %s", val+1, slidesHeaderText[val])
			maxLineLen = max(maxLineLen, len(line))
		}
		lines = append(lines, line)
	}

	maxLineLen = max(maxLineLen, m.TagBar.Width)
	var tags []string
	for _, line := range lines {
		tag := fmt.Sprintf("%-*s\n", maxLineLen, line)
		tags = append(tags, tag)
	}
	return lipgloss.JoinVertical(lipgloss.Left, tags...)
}

func (m *Model) joinHorizontal(width int, cols ...string) string {
	var length int
	for _, col := range cols {
		length += lipgloss.Width(col)
	}

	if width < length {
		return strings.Join(cols, " ")
	}

	var gapCount int
	if len(cols) == 1 {
		gapCount = 1
	} else {
		gapCount = len(cols) - 1
	}
	sub := width - length
	return strings.Join(cols, strings.Repeat(" ", sub/gapCount))
}

func (m *Model) joinVertical(height int, rows ...string) string {
	var h int
	for _, row := range rows {
		h += lipgloss.Height(row)
	}

	if height < h {
		return strings.Join(rows, "\n")
	}

	var gapCount int
	if len(rows) == 1 {
		gapCount = 1
	} else {
		gapCount = len(rows) - 1
	}
	sub := height - h
	return strings.Join(rows, strings.Repeat("\n", sub/gapCount))
}
