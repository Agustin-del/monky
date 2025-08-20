package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
			&ReturnStatement{
				Token: token.Token{Type: token.RETURN, Literal: "return"},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "x"},
					Value: "x",
				},
			},
			&IfStatement{
				Token: token.Token{Type: token.IF, Literal: "if"},
				Condition: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "z"},
					Value: "z",
				},
				Consequence: &BlockStatement{
					Token: token.Token{Type: token.LBRACE, Literal: "{"},
					Statements: []Statement{
						&ReturnStatement{
							Token: token.Token{Type: token.RETURN, Literal: "return"},
							Value: &Identifier{
								Token: token.Token{Type: token.IDENT, Literal: "y"},
								Value: "y",
							},
						},
					},
				},
				Alternative: &BlockStatement{
					Token: token.Token{Type: token.LBRACE, Literal: "{"},
					Statements: []Statement{
						&LetStatement{
							Token: token.Token{Type: token.LET, Literal: "let"},
							Name: &Identifier{
								Token: token.Token{Type: token.IDENT, Literal: "y"},
								Value: "y",
							},
							Value: &Identifier{
								Token: token.Token{Type: token.IDENT, Literal: "y"},
								Value: "y",
							},
						},
					},
				},
			},
		},
	}
	if program.String() != "let myVar = anotherVar;return x;if z {return y;} else {let y = y;}" {
		t.Errorf("program.String() wrong got:%q", program.String())
	}
}
