package model

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/0x00-ketsu/smooth/term/styles"
	"github.com/charmbracelet/lipgloss"
)

// Render draws layout to a string.
func (m *Model) Render() string {
	scrollPercent := m.viewport.ScrollPercent() * 100
	if math.IsNaN(scrollPercent) {
		scrollPercent = 0
	}
	status := fmt.Sprintf("scroll %3.f%%", scrollPercent)

	var main string
	if m.activeTagbar {
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

	var bottom string
	helpView := m.help.View(m.keys)
	if m.Search.Active {
		bottom = m.Search.SearchTextInput.View()
	} else {
		bottom = joinHorizontal(
			m.viewport.Width,
			helpView,
			styles.Status.Render(status),
			styles.Page.Render(fmt.Sprintf("slide %d / %d", m.CurrentPage()+1, m.PageSize())),
		)
	}

	layout := joinVertical(m.viewport.Height+2, main, bottom)

	return layout
}

// render tagbar style string
func (m *Model) renderTagbar() string {
	var tagbar string
	var renderTags []string

	allSlidesHeaderText := m.getAllSlidesHeaderText()
	keys := make([]int, 0, len(allSlidesHeaderText))

	// sort keys desc
	for k := range allSlidesHeaderText {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// fill tagbar
	var line string
	for _, val := range keys {
		if val == m.CurrentPage() {
			line = fmt.Sprintf("⌖ Slide %d: %s\n", val+1, allSlidesHeaderText[val])
			line = styles.ActiveTagbar.Render(line)
		} else {
			line = fmt.Sprintf("Slide %d: %s\n", val+1, allSlidesHeaderText[val])
		}
		renderTags = append(renderTags, line)
	}
	tagbar = lipgloss.JoinVertical(lipgloss.Left, renderTags...)

	return tagbar
}

func joinHorizontal(width int, cols ...string) string {
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

func joinVertical(height int, rows ...string) string {
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
