package ast

import (
	"monkey/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

type IfStatement struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

type Identifier struct {
	Token token.Token
	Value string
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out strings.Builder

	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
	}

	return out.String()
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) String() string {
	var out strings.Builder

	out.WriteString(ls.TokenLiteral() + " ")

	out.WriteString(ls.Name.String())

	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) String() string {
	var out strings.Builder

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.Value != nil {
		out.WriteString(rs.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (is *IfStatement) statementNode() {}

func (is *IfStatement) String() string {
	var out strings.Builder
	out.WriteString(is.TokenLiteral() + " ")
	if is.Condition != nil {
		out.WriteString(is.Condition.String())
	}

	if is.Consequence != nil {
		out.WriteString(" {")
		out.WriteString(is.Consequence.String())
		out.WriteString("}")
	}

	if is.Alternative != nil {
		out.WriteString(" else {")
		out.WriteString(is.Alternative.String())
		out.WriteString("}")
	}

	return out.String()
}

func (is *IfStatement) TokenLiteral() string {
	return is.Token.Literal
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) String() string {
	var out strings.Builder
	for _, stmt := range bs.Statements {
		out.WriteString(stmt.String())
	}

	return out.String()
}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}
