// Package parser contains the parser for parsing tarka tokens and creating an abstract syntax tree.
package parser

import (
	"strconv"

	"github.com/adinovap20/tarkac/internal/ast"
	"github.com/adinovap20/tarkac/internal/token"
)

// Parser structure
type Parser struct {
	tokens    []token.Token // List of tokens using which we need to generate abstract syntax tree
	tokensLen int           // Size of tokens
	curPos    int           // Position of the current pointer
	curToken  token.Token   // Token on which the current pointer is pointing
	errors    []string      // Contains all the parsing errors
}

// NewParser creates a new instance of the Parser structure and returns it.
func NewParser(tokens []token.Token) *Parser {
	p := &Parser{tokens: tokens, tokensLen: len(tokens), curPos: -1}
	p.nextToken()
	return p
}

// Parse returns the abstract syntax tree of the program and returns the Program structure
func (p *Parser) Parse() *ast.Program {
	stmts := []ast.Statement{}
	for p.curPos < p.tokensLen {
		if p.curToken.Type == token.EX_NEWLINE {
			p.nextToken()
			continue
		}
		stmt := p.parseStatement()
		if stmt != nil {
			stmts = append(stmts, stmt)
		} else {
			p.nextToken()
		}
	}
	return &ast.Program{Stmts: stmts}
}

// nextToken moves the current pointer forward
func (p *Parser) nextToken() {
	p.curPos++
	if p.curPos >= p.tokensLen {
		return
	}
	p.curToken = p.tokens[p.curPos]
}

// expectAndConsume expects the token with the given token type at the current position.
// If the expected token type and current token type does not match, it adds parser error in the errors slice.
// This function also moves the pointer forward anyway.
func (p *Parser) expectAndConsume(tokenType token.TokenType) {
	if p.curPos >= p.tokensLen {
		err := "Expected " + string(tokenType) + ", found EOF instead."
		p.newErrorWithoutLineStats(err)
		return
	}
	if p.curToken.Type != tokenType {
		err := "Expected " + string(tokenType) + ", found '" + p.curToken.Lit + "' instead, "
		err += "at line " + strconv.Itoa(p.curToken.Line) + ":" + strconv.Itoa(p.curToken.Col) + "\n"
		p.newError(err, p.curToken.Line, p.curToken.Col)
	}
	p.nextToken()
}
