package model

import (
	"bufio"
	"errors"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/0x00-ketsu/slides/utils"
	"github.com/charmbracelet/glamour"
)

// convert string to markdown style string
func (m *Model) convertToMarkdown(slide string) (string, error) {
	r, _ := glamour.NewTermRenderer(m.termRenderer, glamour.WithWordWrap(m.viewport.Width))

	return r.Render(slide)
}

// extrac slides from markdown file
func (m *Model) extracSlides(content string) []string {
	content = strings.TrimPrefix(content, strings.TrimPrefix(SEP, "\n"))
	slides := strings.Split(content, SEP)

	return slides
}

// get markdown content from file
func (m *Model) readFile() (string, error) {
	if !utils.IsFileExist(m.Filename) {
		return "", errors.New("File is not exist")
	}

	content, err := os.ReadFile(m.Filename)
	return string(content), err
}

// get markdown content from stdin
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

// extract markdown header text from every single slide, iterate all slides.
// If not find markdown header return the first unempty line (part)
func (m *Model) getAllSlidesHeaderText() map[int]string {
	store := make(map[int]string)
	for idx, slide := range m.Slides {
		headerText := extractHeaderText(slide)
		if headerText == "" {
			unemptyLine := getFirstUnemptyLine(slide)
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

// return first unempty(stripped) line of plain slide
func getFirstUnemptyLine(plainSlide string) string {
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

// extract markdown header text if matched in slide
// else return empty string
func extractHeaderText(plainSlide string) string {
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
