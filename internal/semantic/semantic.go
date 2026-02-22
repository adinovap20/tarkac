// Package semantic contains the semantic analyzer for the tarka abstract syntax tree
package semantic

import (
	"strconv"

	"github.com/adinovap20/tarkac/internal/ast"
)

// SemanticAnalyzer structure
type SemanticAnalyzer struct {
	errors    []string
	lastType  string
	variables map[string]Type
}

// NewSemanticAnalyzer creates a new instance of SemanticAnalyzer structure
func NewSemanticAnalyzer() *SemanticAnalyzer {
	return &SemanticAnalyzer{
		errors:    []string{},
		variables: map[string]Type{},
		lastType:  "",
	}
}

// VisitProgram analyzes the Program node
func (s *SemanticAnalyzer) VisitProgram(program *ast.Program) {
	for _, stmt := range program.Stmts {
		stmt.Accept(s)
	}
}

// VisitStmtExit analyzes the StmtExit node
func (s *SemanticAnalyzer) VisitStmtExit(stmtExit *ast.StmtExit) {
	if stmtExit.Expr == nil {
		err := "Exit statement expected an expression"
		s.newError(err, stmtExit.Line, stmtExit.Col)
		return
	}
	stmtExit.Expr.Accept(s)
	if s.lastType != TYPE_INT {
		err := "Exit statement expected int type, got " + s.lastType + " instead"
		s.newError(err, stmtExit.Line, stmtExit.Col)
		return
	}
	if exprIntLit, ok := stmtExit.Expr.(*ast.ExprIntLit); ok {
		if exprIntLit.Val < 0 || exprIntLit.Val > 255 {
			err := "Exit statement expected a value between 0 and 255, got " + strconv.Itoa(exprIntLit.Val) + " instead"
			s.newError(err, exprIntLit.Line, exprIntLit.Col)
			return
		}
	}
}

// VisitStmtIntVarDecl analyzes the StmtIntVarDecl node
func (s *SemanticAnalyzer) VisitStmtIntVarDecl(stmtIntVarDecl *ast.StmtIntVarDecl) {
	if _, ok := s.variables[stmtIntVarDecl.Name]; ok {
		err := "Variable " + stmtIntVarDecl.Name + " already declared"
		s.newError(err, stmtIntVarDecl.Line, stmtIntVarDecl.Col)
		return
	}
	s.variables[stmtIntVarDecl.Name] = TYPE_INT
}

// VisitExprIntLit analyzes the ExprIntLit node
func (s *SemanticAnalyzer) VisitExprIntLit(exprIntLit *ast.ExprIntLit) {
	s.lastType = TYPE_INT
}
