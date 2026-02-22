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
	case token.LIT_IDENT:
		if p.peekToken(1).Type == token.PUNC_COLON && p.peekToken(2).Type == token.KW_INT {
			return p.parseStmtIntVarDecl()
		}
		p.newNoMatchingStmtError()
		return nil
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

// parseStmtIntVarDecl parses the int variable declaration statement and returns the StmtIntVarDecl node
func (p *Parser) parseStmtIntVarDecl() *ast.StmtIntVarDecl {
	line, col := p.curToken.Line, p.curToken.Col
	varName := p.curToken.Lit
	p.expectAndConsume(token.LIT_IDENT)
	p.expectAndConsume(token.PUNC_COLON)
	p.expectAndConsume(token.KW_INT)
	p.expectAndConsume(token.OP_ASSIGN)
	expr := p.parseExpression()
	if expr == nil {
		p.newError("Invalid expression after variable declaration", line, col)
		return nil
	}
	p.expectAndConsume(token.EX_NEWLINE)
	return &ast.StmtIntVarDecl{Name: varName, Expr: expr, Line: line, Col: col}
}
