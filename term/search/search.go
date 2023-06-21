package search

import (
	"regexp"
	"strings"

	"github.com/0x00-ketsu/slides/term/styles"
	"github.com/charmbracelet/bubbles/textinput"
)

type Model interface {
	GetAllSlides() []string
	CurrentPage() int
	PageSize() int
	SetPage(page int)
}

type Search struct {
	Active          bool
	SearchTextInput textinput.Model
	// Store searched text slides
	Cache             []int
	CurrentCacheIndex int
}

func NewSearch() Search {
	tm := textinput.New()
	tm.Placeholder = "search ..."
	tm.Prompt = "/"
	tm.PromptStyle = styles.Search
	tm.TextStyle = styles.Search

	return Search{SearchTextInput: tm}
}

// Begin starts a new search (deltes old buffers)
func (s *Search) Begin() {
	s.Active = true
	s.SetQuery("")
	s.CurrentCacheIndex = 0
	s.Cache = []int{}
}

func (s *Search) Query() string {
	return s.SearchTextInput.Value()
}

func (s *Search) SetQuery(query string) {
	s.SearchTextInput.SetValue(query)
}

func (s *Search) Next(m Model) {
	caches := s.Cache
	if len(caches) > 0 {
		s.CurrentCacheIndex++
		if s.CurrentCacheIndex == len(caches) {
			s.CurrentCacheIndex = 0
		}
		m.SetPage(caches[s.CurrentCacheIndex])
	}
}

func (s *Search) Previous(m Model) {
	caches := s.Cache
	if len(caches) > 0 {
		s.CurrentCacheIndex--
		if s.CurrentCacheIndex == -1 {
			s.CurrentCacheIndex = len(caches) - 1
		}
		m.SetPage(caches[s.CurrentCacheIndex])
	}
}

func (s *Search) Done() {
	s.Active = false
}

func (s *Search) Execute(m Model) {
	defer s.Done()

	expr := s.Query()
	if expr == "" {
		return
	}

	if strings.HasPrefix(expr, "/i") {
		expr = "(?i)" + expr[:len(expr)-2]
	}

	pattern, err := regexp.Compile(expr)
	if err != nil {
		return
	}
	slides := m.GetAllSlides()
	addCache := func(i int) {
		content := slides[i]
		if len(pattern.FindAllStringSubmatch(content, 1)) != 0 {
			if !s.isCached(i) {
				s.Cache = append(s.Cache, i)
			}
		}
	}

	// forward search
	// rules:
	//   1. current slide -> last slide
	//   2. first slide -> current slide
	// Search from current slide to end
	for i := m.CurrentPage(); i < len(slides); i++ {
		addCache(i)
	}

	// Search from first slide to current
	for i := 0; i < m.CurrentPage(); i++ {
		addCache(i)
	}

	// goto first searched slide
	if len(s.Cache) < 1 {
		return
	}
	slideIdx := s.Cache[0]
	s.CurrentCacheIndex = 0
	m.SetPage(slideIdx)
}

func (s *Search) isCached(target int) bool {
	if len(s.Cache) < 1 {
		return false
	}

	for _, val := range s.Cache {
		if target == val {
			return true
		}
	}

	return false
}
