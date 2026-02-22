// Package token contains all the information related to tarka tokens
package token

import (
	"fmt"
	"strings"
)

// TokenType enum contains all the different types of tokens
type TokenType string

const (
	KW_EXIT = "KW_EXIT" // KW_EXIT is token type for the 'exit' keyword
	KW_INT  = "KW_INT"  // KW_INT is token type for the 'int' keyword

	LIT_INT   = "LIT_INT"   // LIT_INT is token type for the integer literal, e.g., 10, 302, etc.
	LIT_IDENT = "LIT_IDENT" // LIT_IDENT is token type for the identifier literal, e.g., a, b, etc.

	OP_ASSIGN = "OP_ASSIGN" // OP_ASSIGN is token type for '='

	PUNC_COLON = "PUNC_COLON" // PUNC_COLON is token type for ':'

	EX_UNKNOWN = "EX_UNKNOWN" // EX_UNKNOWN is the token type for unknown token
	EX_NEWLINE = "EX_NEWLINE" // EX_NEWLINE is the token type for newline token
)

// Token structure
type Token struct {
	Type TokenType // Type stores the type of the token
	Lit  string    // Lit stores the literal associated with the token
	Line int       // Line stores the line number associated with the beginning of the token
	Col  int       // Col stores the column number associated with the beginning of the token
}

// KEYWORDS map stores keyword literals and their corresponding keyword token type
var KEYWORDS = map[string]TokenType{
	"exit": KW_EXIT,
	"int":  KW_INT,
}

// LookupIdent looks up identifier in the KEYWORDS map and return the corresponding token type.
// This returns LIT_IDENT token type if the provided identifier literal is not a keyword.
func LookupIdent(ident string) TokenType {
	if tokType, ok := KEYWORDS[ident]; ok {
		return tokType
	}
	return LIT_IDENT
}

// PrintTokens prints the tokens in a nice way
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
