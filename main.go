package main 

import (
	"io/ioutil"
	"fmt"
	"github.com/cloudson/lingo/parser"
	"github.com/cloudson/lingo/symbol"
)

func main() {
	example := "claudson"
	f, _:= ioutil.ReadFile("./fonts/1.lf")
	p := parser.New(string(f))
	for i := 1; i <= 3; i++ {
		for _, c := range example {
			char, error := p.Char(c)
			if error != nil {
				panic(error)
			}
			str, _:= symbol.Print(char, i)
			fmt.Printf("%s", str)
			// fmt.Printf("%s\n", char)	
		}
		fmt.Print("\n")	
	}
	
}