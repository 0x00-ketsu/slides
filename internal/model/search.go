package model

import (
	"regexp"
	"strings"

	"github.com/0x00-ketsu/slides/internal/styles"
	"github.com/charmbracelet/bubbles/textinput"
)

// Search is the search model.
type Search struct {
	// IsActive indicates if search is active
	IsActive bool

	// TextInput is the search input
	TextInput textinput.Model

	// SearchResultsCache is the cache of search results
	SearchResultsCache []int

	// CurrentCachedIndex is the current index of the cache
	CurrentCachedIndex int
}

// NewSearch creates a new search model.
func NewSearch() *Search {
	input := textinput.New()
	input.Placeholder = "input search text"
	input.Prompt = "/"
	input.PromptStyle = styles.Search
	input.TextStyle = styles.Search
	return &Search{TextInput: input}
}

// Begin starts a new search (delete old buffers)
func (s *Search) Begin() {
	s.IsActive = true
	s.SetQuery("")
	s.CurrentCachedIndex = 0
	s.SearchResultsCache = []int{}
}

func (s *Search) Query() string {
	return s.TextInput.Value()
}

func (s *Search) SetQuery(query string) {
	s.TextInput.SetValue(query)
}

func (s *Search) Next(m *Model) {
	cache := s.SearchResultsCache
	cacheSize := len(cache)
	if cacheSize > 0 {
		s.CurrentCachedIndex++
		if s.CurrentCachedIndex == cacheSize {
			s.CurrentCachedIndex = 0
		}
		m.SetPage(cache[s.CurrentCachedIndex])
	}
}

func (s *Search) Previous(m *Model) {
	cache := s.SearchResultsCache
	if len(cache) > 0 {
		s.CurrentCachedIndex--
		if s.CurrentCachedIndex == -1 {
			s.CurrentCachedIndex = len(cache) - 1
		}
		m.SetPage(cache[s.CurrentCachedIndex])
	}
}

func (s *Search) Done() {
	s.IsActive = false
}

func (s *Search) Execute(m *Model) {
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
				s.SearchResultsCache = append(s.SearchResultsCache, i)
			}
		}
	}

	// forward search
	// rules:
	//   1. current slide -> last slide
	//   2. first slide -> current slide
	// Search from current slide to end
	for i := m.GetCurrentPage(); i < len(slides); i++ {
		addCache(i)
	}

	// Search from first slide to current
	for i := 0; i < m.GetCurrentPage(); i++ {
		addCache(i)
	}

	// goto first searched slide
	if len(s.SearchResultsCache) < 1 {
		return
	}
	slideIdx := s.SearchResultsCache[0]
	s.CurrentCachedIndex = 0
	m.SetPage(slideIdx)
}

func (s *Search) isCached(target int) bool {
	if len(s.SearchResultsCache) < 1 {
		return false
	}

	for _, val := range s.SearchResultsCache {
		if target == val {
			return true
		}
	}

	return false
}
