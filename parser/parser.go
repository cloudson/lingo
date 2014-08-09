package parser

import (
	"fmt"
	"strings"
	"strconv"
	"errors"	
	"github.com/cloudson/lingo/symbol"
)

type Parser struct {
	Header *Header
	contentRaw string
	body string
}

type Header struct {
	name string
	Height int 
	numberOfLines int 
	positions map[rune]int
	alphabet string
}

func New(fileContent string) (*Parser) {
	p := new(Parser)
	p.contentRaw = fileContent
	var e error
	p.Header, e = generateHeader(p.contentRaw)
	if e != nil {
		panic(e)
	}

	lines := strings.Split(p.contentRaw, "\n")
	
	if len(lines) > p.Header.numberOfLines + 1 {
		p.body = strings.Join(lines[p.Header.numberOfLines +1:], "\n")
	}	

	return p
}

func (p *Parser) getPositions() (map[rune]int) {
	var positions = make(map[rune]int)
	for pos, char := range p.Header.alphabet {
		positions[char] = pos + 1
	}

	return positions
}

func (p *Parser) Char(char rune) (*symbol.Symbol, error) {
	positions := p.getPositions()
	position, ok := positions[char]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Char '%s' is not found", string(char)))
	}
	
	contentFiles := strings.Split(p.body, "\n")
	if p.body == ""{
		return nil, errors.New("Char has not ascii code") 
	}
	charRange := contentFiles[(position-1) * p.Header.Height:(position -1)* p.Header.Height + p.Header.Height]
	charResult := strings.Join(charRange, "\n")

	return symbol.New(charResult), nil
}


func generateHeader(content string) (*Header, error) {
	Header := new(Header)
	Header.numberOfLines = 0
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		metadataLine := strings.Split(line, ":")
		if len(metadataLine) >= 2 {
			value := strings.Trim(metadataLine[1], " ")
			switch strings.Trim(metadataLine[0], " ") {
				case "Name": 
					Header.name = value
				case "Height" :
					Header.Height, _ = strconv.Atoi(value)
				case "Alphabet":
					Header.alphabet = value

			}
		}
		if strings.Index(line, "===") == 0 {
			break
		}
		Header.numberOfLines += 1
	}
	var e error
	if Header.name == "" {
		e = errors.New("Field 'name' not declared on header")
	}
	if Header.Height == 0 {
		e = errors.New("Field 'height' not declared on header")
	}
	if Header.alphabet == "" {
		e = errors.New("Field 'alphabet' not declared on header")
	}

	return Header,e
}
