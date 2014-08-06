package symbol 

import (
	"strings"
	"unicode/utf8"
)

type Symbol struct {
	raw string
	lines []string
	maxWidth int 
}

func New(char string) (*Symbol) {
	s := new(Symbol)
	s.raw = char
	s.maxWidth = s.width()

	return s
}

func (s *Symbol) width() int {
	lines := strings.Split(s.raw, "\n")
	maxWidth := 0 
	for _, line := range lines {
		length := utf8.RuneCountInString(strings.Trim(line, " "))
		if length > maxWidth {
			maxWidth = length
		}
	}

	return maxWidth
}

func (s *Symbol) String() string {
	return  s.raw
}
