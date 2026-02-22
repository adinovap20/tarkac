// Package lexer contains the tarka lexer
package lexer

import (
	"github.com/adinovap20/tarkac/internal/token"
	"github.com/adinovap20/tarkac/internal/utils"
)

// Lexer structure
type Lexer struct {
	code    string // Code string to tokenize
	codeLen int    // Size of the code
	curPos  int    // Location of the pointer
	curChar byte   // Character at the current pointer
	line    int    // Line number of the character pointed by current pointer
	col     int    // Column number of the character pointed by current pointer
}

// NewLexer creates a new instance of Lexer and returns it
func NewLexer(code string) *Lexer {
	lexer := &Lexer{
		code:    code,
		codeLen: len(code),
		curPos:  -1,
		curChar: '-',
		line:    1,
		col:     0,
	}
	lexer.readChar()
	return lexer
}

// GetTokens tokenizes the code string using which the Token structure was initialized
func (l *Lexer) GetTokens() []token.Token {
	tokens := []token.Token{}
	for l.curPos < l.codeLen {
		if l.curChar == ' ' || l.curChar == '\t' || l.curChar == '\r' {
			l.readChar()
			continue
		}
		if l.curChar == '\n' {
			tok := l.readNewline()
			tokens = append(tokens, tok)
			continue
		}
		if utils.IsLetter(l.curChar) {
			tok := l.readIdentifier()
			tokens = append(tokens, tok)
			continue
		}
		if utils.IsDigit(l.curChar) {
			tok := l.readNumber()
			tokens = append(tokens, tok)
			continue
		}
		if l.curChar == '=' {
			tok := l.readAssignment()
			tokens = append(tokens, tok)
			continue
		}
		if l.curChar == ':' {
			tok := l.readColon()
			tokens = append(tokens, tok)
			continue
		}
		tok := l.readUnknown()
		tokens = append(tokens, tok)
	}
	return tokens
}

// readChar moves the current pointer forward. If the current pointer moves ahead of the length of
// the code, it returns. If a newline character is found, it increases the current line number by 1 and
// makes the current column number 0.
func (l *Lexer) readChar() {
	l.curPos++
	if l.curPos >= l.codeLen {
		return
	}
	if l.curChar == '\n' {
		l.line++
		l.col = 0
	}
	l.curChar = l.code[l.curPos]
	l.col++
}

// readIdentifier reads the identifier and returns either identifier or a keyword token
func (l *Lexer) readIdentifier() token.Token {
	beg, line, col := l.curPos, l.line, l.col
	for utils.IsAlphaNumeric(l.curChar) && l.curPos < l.codeLen {
		l.readChar()
	}
	lit := l.code[beg:l.curPos]
	tokType := token.LookupIdent(lit)
	tok := token.Token{Type: tokType, Lit: lit, Line: line, Col: col}
	return tok
}

// readNumber reads the number and returns the corresponding token
func (l *Lexer) readNumber() token.Token {
	beg, line, col := l.curPos, l.line, l.col
	for utils.IsDigit(l.curChar) && l.curPos < l.codeLen {
		l.readChar()
	}
	lit := l.code[beg:l.curPos]
	tok := token.Token{Type: token.LIT_INT, Lit: lit, Line: line, Col: col}
	return tok
}

// readAssignment reads the assignment operator and returns the corresponding token
func (l *Lexer) readAssignment() token.Token {
	tok := token.Token{Type: token.OP_ASSIGN, Lit: "=", Line: l.line, Col: l.col}
	l.readChar()
	return tok
}

// readColon reads the colon and returns the corresponding token
func (l *Lexer) readColon() token.Token {
	tok := token.Token{Type: token.PUNC_COLON, Lit: ":", Line: l.line, Col: l.col}
	l.readChar()
	return tok
}

// readNewline reads the current character and returns the newline token
func (l *Lexer) readNewline() token.Token {
	tok := token.Token{Type: token.EX_NEWLINE, Lit: "\n", Line: l.line, Col: l.col}
	l.readChar()
	return tok
}

// readUnknown reads the current character as unknown token and returns the newly created token
func (l *Lexer) readUnknown() token.Token {
	tok := token.Token{Type: token.EX_UNKNOWN, Lit: string(l.curChar), Line: l.line, Col: l.col}
	l.readChar()
	return tok
}
