package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	curToken  token.Token
	peekToken token.Token
	lexer     *lexer.Lexer
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l, errors: []string{}}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.IF:
		return p.parseIfStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		p.errors = append(p.errors, "Se esperaba un identificador.")
		p.skipTilSemicolon()
		return nil
	}

	statement.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		p.errors = append(p.errors, "Se esperaba '='")
		p.skipTilSemicolon()
		return nil
	}

	p.nextToken()

	statement.Value = p.parseExpression()

	if !p.expectPeek(token.SEMICOLON) {
		p.errors = append(p.errors, "Se esperaba ';'")
		p.skipTilSemicolon()
		return nil
	}
	return statement
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: p.curToken}

	if p.curToken.Type != token.RETURN {
		p.errors = append(p.errors, "Se esperaba 'return'")
		p.skipTilSemicolon()
		return nil
	}

	p.nextToken()
	statement.Value = p.parseExpression()

	p.nextToken()
	if p.curToken.Type != token.SEMICOLON {
		p.errors = append(p.errors, "Se esperaba ';'")
		return nil
	}

	return statement
}

func (p *Parser) parseIfStatement() *ast.IfStatement {
	return nil
}

func (p *Parser) parseExpression() ast.Expression {
	return nil
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if t == p.peekToken.Type  {
		p.nextToken()
		return true
	} else {
		p.peekError(p.peekToken.Type)
		return false
	}
}

func (p *Parser) peekError(t token.TokenType){
	msg := fmt.Sprintf("Expected next token to be %s, got %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) skipTilSemicolon() {
	for p.curToken.Type != token.SEMICOLON && p.curToken.Type != token.EOF {
		p.nextToken()
	}

	if p.curToken.Type == token.SEMICOLON {
		p.nextToken()
	}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}
