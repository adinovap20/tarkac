package parser

import (
	"reflect"
	"testing"

	"github.com/adinovap20/tarkac/internal/ast"
	"github.com/adinovap20/tarkac/internal/lexer"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected ast.Program
	}{
		{"exit 10\nexit 20\n", ast.Program{
			Stmts: []ast.Statement{
				&ast.StmtExit{Expr: &ast.ExprIntLit{Val: 10, Line: 1, Col: 6}, Line: 1, Col: 1},
				&ast.StmtExit{Expr: &ast.ExprIntLit{Val: 20, Line: 2, Col: 6}, Line: 2, Col: 1},
			},
		}},
		{"a: int = 10\n", ast.Program{
			Stmts: []ast.Statement{
				&ast.StmtIntVarDecl{Expr: &ast.ExprIntLit{Val: 10, Line: 1, Col: 10}, Line: 1, Col: 1},
			},
		}},
	}
	for _, tt := range tests {
		lexer := lexer.NewLexer(tt.input)
		toks := lexer.GetTokens()
		parser := NewParser(toks)
		program := parser.Parse()
		if program == nil {
			t.Errorf("Expected program, got nil")
			continue
		}
		if len(program.Stmts) != len(tt.expected.Stmts) {
			t.Errorf("Expected %d statements, got %d", len(tt.expected.Stmts), len(program.Stmts))
			continue
		}
		for i, stmt := range program.Stmts {
			if !reflect.DeepEqual(stmt, tt.expected.Stmts[i]) {
				t.Errorf("stmt %d wrong content. got=%+v, want=%+v", i, stmt, tt.expected.Stmts[i])
				continue
			}
		}
	}
}
