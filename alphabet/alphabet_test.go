package alphabet

import (
	"testing"
)

func TestToString(t *testing.T) {

	expected := `┌┬┐
│││
┴ ┴
`
	symbol := CreateSymbol(expected)
	if symbol.String() != expected {
		t.Errorf("Expected:\n'%s'\nFound:\n'%s'\n", expected, symbol.String())
	}
}

func TestMaxWidth(t *testing.T) {
	char := `┌─┐
│  
└─┘
`
	expected := 3
	symbol := CreateSymbol(char)
	if expected != symbol.maxWidth {
		t.Errorf("Max width expected %d, found %d", expected, symbol.maxWidth)
	}
}