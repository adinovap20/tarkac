package parser

import (
	"github.com/adinovap20/tarkac/internal/ast"
	"github.com/adinovap20/tarkac/internal/token"
)

// parseStatement parses the statement and returns the statement node
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.KW_EXIT:
		return p.parseStmtExit()
	}
	p.newNoMatchingStmtError()
	return nil
}

// parseStmtExit parses the exit statement and returns the StmtExit node
func (p *Parser) parseStmtExit() *ast.StmtExit {
	line, col := p.curToken.Line, p.curToken.Col
	p.nextToken() // exit keyword parsed
	expr := p.parseExpression()
	if expr == nil {
		p.newError("Invalid expression after exit statement", line, col)
		return nil
	}
	p.expectAndConsume(token.EX_NEWLINE)
	return &ast.StmtExit{Expr: expr, Line: line, Col: col}
}
