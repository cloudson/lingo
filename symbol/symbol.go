package symbol 

import (
	"errors"
	"strings"
	"unicode/utf8"
)

type Symbol struct {
	raw string
	lines []string
}

func New(char string) (*Symbol) {
	s := new(Symbol)
	s.raw = char

	return s
}

func (s *Symbol) Width() int {
	lines := strings.Split(s.raw, "\n")
	maxWidth := 0 
	for _, line := range lines {
		length := utf8.RuneCountInString(strings.TrimRight(line, " "))
		if length > maxWidth {
			maxWidth = length
		}
	}

	return maxWidth
}

func (s *Symbol) String() string {
	return  s.raw
}


func Print(s *Symbol, position int) (string, error) {
	lines := strings.Split(s.raw, "\n")

	if  position > len(lines) - 1 {
		return "", errors.New("Trying access inexisting line")
	}

	lineRaw := strings.TrimRight(lines[position-1], " ")
	lineResult := lineRaw + strings.Repeat(" ", s.Width() - utf8.RuneCountInString(strings.TrimRight(lineRaw, " ")))


	return lineResult, nil
}