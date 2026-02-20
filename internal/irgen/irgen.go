// Package irgen contains visitor for generating IR
package irgen

import (
	"github.com/adinovap20/tarkac/internal/ast"
	"github.com/adinovap20/tarkac/internal/ir"
)

// IRGenerator visitor structure that visitos AST and creates IR
type IRGenerator struct {
	IRProgram *ir.IRProgram
}

// NewIRGenerator creates and returns a new instance of IRGenerator
func NewIRGenerator() *IRGenerator {
	return &IRGenerator{IRProgram: ir.New()}
}

// Print prints the IRProgram
func (g *IRGenerator) Print() {
	println(g.IRProgram.String())
}

// VisitProgram visits the program node and generates corresponding IR
func (g *IRGenerator) VisitProgram(program *ast.Program) {
	for _, stmt := range program.Stmts {
		if stmt != nil {
			stmt.Accept(g)
		}
	}
}

// VisitStmtExit visits the exit statement
func (g *IRGenerator) VisitStmtExit(stmtExit *ast.StmtExit) {
	if stmtExit.Expr != nil {
		stmtExit.Expr.Accept(g)
	}
	g.IRProgram.Exit()
}

// VisitExprIntLit visits the integer literal expression
func (g *IRGenerator) VisitExprIntLit(exprIntLit *ast.ExprIntLit) {
	g.IRProgram.LoadInt(exprIntLit.Val)
}
