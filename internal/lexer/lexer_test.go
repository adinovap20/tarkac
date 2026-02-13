package lexer

import (
	"testing"

	"github.com/adinovap20/tarkac/internal/token"
)

// TestLexer tests lexer in an end to end way. It checks if tokens are getting properly generated for the
// provided code string. If they are not generating, the tests will fail.
func TestLexer(t *testing.T) {
	tests := []struct {
		input    string
		expected []token.Token
	}{
		{"exit 10\nhello", []token.Token{
			{Type: token.KW_EXIT, Lit: "exit", Line: 1, Col: 1},
			{Type: token.LIT_INT, Lit: "10", Line: 1, Col: 6},
			{Type: token.EX_NEWLINE, Lit: "\n", Line: 1, Col: 8},
			{Type: token.LIT_IDENT, Lit: "hello", Line: 2, Col: 1},
		}},
	}
	for _, tt := range tests {
		lexer := NewLexer(tt.input)
		toks := lexer.GetTokens()
		if len(toks) != len(tt.expected) {
			t.Errorf("Expected %d tokens, got %d", len(tt.expected), len(toks))
			continue
		}
		for i, tok := range toks {
			if tok.Type != tt.expected[i].Type {
				t.Errorf("Expected token type %s, got %s", tt.expected[i].Type, tok.Type)
			}
			if tok.Lit != tt.expected[i].Lit {
				t.Errorf("Expected token literal %s, got %s", tt.expected[i].Lit, tok.Lit)
			}
			if tok.Line != tt.expected[i].Line {
				t.Errorf("Expected token line %d, got %d", tt.expected[i].Line, tok.Line)
			}
			if tok.Col != tt.expected[i].Col {
				t.Errorf("Expected token column %d, got %d", tt.expected[i].Col, tok.Col)
			}
		}
	}
}
