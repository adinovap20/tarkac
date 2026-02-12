// token package contains all the information related to tarka tokens
package token

import (
	"fmt"
	"strings"
)

// TokenType enum contains all the different types of tokens
type TokenType string

const (
	KW_EXIT = "KW_EXIT" // exit keyword

	LIT_INT   = "LIT_INT"   // integer literal, e.g., 10, 302, etc.
	LIT_IDENT = "LIT_IDENT" // identifier literal, e.g., a, b, etc.

	EX_UNKNOWN = "EX_UNKNOWN" // unknown token
	EX_NEWLINE = "EX_NEWLINE" // newline token
)

// token structure
type Token struct {
	Type TokenType // type of the token
	Lit  string    // token literal
	Line int       // line number where token begins
	Col  int       // column number where token begins
}

// keywords map that stores keyword literals and their corresponding keyword token type
var KEYWORDS = map[string]TokenType{
	"exit": KW_EXIT,
}

// function to lookup identifier in the KEYWORDS map and return the corresponding token type.
// This returns LIT_IDENT token type if the provided identifier literal is not a keyword.
func LookupIdent(ident string) TokenType {
	if tokType, ok := KEYWORDS[ident]; ok {
		return tokType
	}
	return LIT_IDENT
}

// function to print tokens in a nice way
func PrintTokens(tokens []Token) {
	for i, tok := range tokens {
		lit := strings.ReplaceAll(tok.Lit, "\n", "\\n")
		fmt.Printf("[%s,'%s',%d:%d]", tok.Type, lit, tok.Line, tok.Col)
		if i < len(tokens)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Println()
}
