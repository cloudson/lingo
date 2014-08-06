package parser

import (
	"fmt"
	"strings"
	"strconv"
	"errors"	
	"github.com/cloudson/lingo/symbol"
)

type Parser struct {
	header *Header
	contentRaw string
	body string
}

type Header struct {
	name string
	height int 
	numberOfLines int 
	positions map[rune]int
	alphabet string
}

func New(fileContent string) (*Parser) {
	p := new(Parser)
	p.contentRaw = fileContent
	p.header = generateHeader(p.contentRaw)
	lines := strings.Split(p.contentRaw, "\n")
	
	if len(lines) > p.header.numberOfLines + 1 {
		p.body = strings.Join(lines[p.header.numberOfLines +1:], "\n")
	}	

	return p
}

func (p *Parser) getPositions() (map[rune]int) {
	var positions = make(map[rune]int)
	for pos, char := range p.header.alphabet {
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
	charRange := contentFiles[(position-1) * p.header.height:(position -1)* p.header.height + p.header.height]
	charResult := strings.Join(charRange, "\n")

	return symbol.New(charResult), nil
}


func generateHeader(content string) *Header {
	header := new(Header)
	header.numberOfLines = 0
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		metadataLine := strings.Split(line, ":")
		if len(metadataLine) >= 2 {
			value := strings.Trim(metadataLine[1], " ")
			switch strings.Trim(metadataLine[0], " ") {
				case "Name": 
					header.name = value
				case "Height" :
					header.height, _ = strconv.Atoi(value)
				case "Alphabet":
					header.alphabet = value

			}
		}
		if strings.Index(line, "===") == 0 {
			break
		}
		header.numberOfLines += 1
	}
	return header
}
