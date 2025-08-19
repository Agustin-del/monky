package main

import (
	"monkey/lexer"
	"monkey/parser"
)

func main() {
	input := `let x = 5;`

	l := lexer.New(input)
	parser.New(l)
}
