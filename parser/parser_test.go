package parser 

import (
	"testing"
	"strings"
)

func TestFontname(t *testing.T) {
	expected := "Cloud Font"
	fileContent := `
Name: Cloud Font 
	`
	h := generateHeader(fileContent)
	found := h.name

	if found != expected {
		t.Errorf("Expected Font name '%s' and found '%s' ", expected, found)
	}
}

func TestFontHeight(t *testing.T) {
	expected := 5
	fileContent := `
Name: Best Font
Height : 5 
	`
	h := generateHeader(fileContent)
	found := h.height

	if found != expected {
		t.Errorf("Expected height %d and found height %d ", expected, found)
	}
}

func TestEndOfHeader(t *testing.T) {
	fileContent := `Name: Best Font
Height : 5 
======
Get: lucky! 
	`
	expected := 2 

	h := generateHeader(fileContent)
	found := h.numberOfLines
	if expected != found {
		t.Errorf("The header should be %d lines instead of %d", expected, found)
	}
}

func TestAValidCharPosition(t *testing.T) {
	fileContent := `Name: Saturno Font
Height: 6
Alphabet: cloud
========
`
	expecteds :=  map[rune]int{'c':1, 'l':2, 'o':3, 'u':4, 'd':5} 

	p := New(fileContent)
	positions := p.getPositions()
	if len(expecteds) != len(positions) {
		t.Errorf("Positions has not the expected length (%d != %d)", len(positions), len(expecteds));
	}
	for char, pos := range positions {
		if expecteds[char] != pos {
			t.Errorf("Symbol '%s' should be at position %d, it's at %d", string(char), pos, expecteds[char])
		}
	}
}

func TestAValidChar(t *testing.T) {
	fileContent := `Name: Saturno Font
Height: 6
Alphabet: c
========
        
  ____  
_/ ___\ 
\  \___ 
 \___  >
     \/`
	expected := 
`        
  ____  
_/ ___\ 
\  \___ 
 \___  >
     \/`

	p := New(fileContent)
	char, error := p.getChar('c');
	if error != nil {
		t.Error(error)
	}
	if char.String() != expected {
		t.Errorf("Expected:\n%s\nFound:\n%s\n", expected, char)
	}
}


func TestValidsChars(t *testing.T) {
	fileContent := `Name: Saturno Font
Height: 6
Alphabet: cl
========
        
  ____  
_/ ___\ 
\  \___ 
 \___  >
     \/
 __ 
|  |  
|  |  
|  |__
|____/
      `
	expected := 
	`
  ____  
_/ ___\ 
\  \___ 
 \___  >
     \/`

	p := New(fileContent)
	char, error := p.getChar('c');
	if error != nil {
		t.Error(error)
	}
	if serialize(char.String()) != serialize(expected) {
		t.Errorf("Expected:\n'%s'\nFound:\n'%s'\n", expected, char)
	}

	expected = ` __ 
|  |  
|  |  
|  |__
|____/
      `
	char1, error1 := p.getChar('l');
	if error1 != nil {
		t.Error(error1)
	}
	if serialize(char1.String()) != serialize(expected) {
		t.Errorf("Expected:\n'%s'\nFound:\n'%s'\n", expected, char1)
	}


}


func serialize(char string) string {
	return strings.Replace(char, " ", "", -1)
}