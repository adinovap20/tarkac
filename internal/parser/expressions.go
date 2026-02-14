package parser

import (
	"strconv"

	"github.com/adinovap20/tarkac/internal/ast"
	"github.com/adinovap20/tarkac/internal/token"
)

// parseExpression tries to parse an expression from the current token and returns an expression node
func (p *Parser) parseExpression() ast.Expression {
	switch p.curToken.Type {
	case token.LIT_INT:
		return p.parseExprIntLit()
	}
	p.newNoMatchingExprError()
	return nil
}

// parseExprIntLit parses the integer literal expression and returns its corresponding node.
func (p *Parser) parseExprIntLit() *ast.ExprIntLit {
	val, err := strconv.Atoi(p.curToken.Lit)
	line, col := p.curToken.Line, p.curToken.Col
	if err != nil {
		p.newError("Invalid integer literal "+p.curToken.Lit+" found", line, col)
		return nil
	}
	p.nextToken()
	return &ast.ExprIntLit{Val: val, Line: line, Col: col}
}
