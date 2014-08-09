package symbol

import (
	"testing"
)

func TestToString(t *testing.T) {

	expected := `┌┬┐
│││
┴ ┴
`
	symbol := New(expected)
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
	symbol := New(char)
	if expected != symbol.Width() {
		t.Errorf("Max width expected %d, found %d", expected, symbol.Width())
	}
}

func TestPrintOneLineWithMaxChars(t *testing.T) {
	char := `┌─┐
│  
└─┘
`
	expected := "└─┘"
	s := New(char)
	found, error := Print(s, 3)
	if error != nil {
		t.Error(error)
	}
	if expected != found {
		t.Errorf("Expected:\n'%s'\nFound:\n'%s'\n", expected, found)
	} 
}

func TestPrintLineOutOfRange(t *testing.T) {
	char := `┌─┐
│  
└─┘`
	
	s:= New(char)
	_, error:= Print(s, 4)
	if error == nil {
		t.Error("Should throw error when print out of range")
	}

	_, error = Print(s, 5)
	if error == nil {
		t.Error("Should throw error when print out of range")
	}

	_, error = Print(s, -1)
	if error == nil {
		t.Error("Should throw error when print out of range")
	}
}

func TestPrintLineAddingSpaces(t *testing.T) {
	char := `┌─┐
│
└─┘
`
	expected := "│  "
	s := New(char)
	found, error := Print(s, 2)
	if error != nil {
		t.Error(error)
	}
	if expected != found {
		t.Errorf("Expected:\n'%s'\nFound:\n'%s'\n", expected, found)
	} 		
}	