// Package astprinter visits the AST and prints it in a nice way.
package astprinter

import "github.com/adinovap20/tarkac/internal/ast"

// ASTPrinter structure
type ASTPrinter struct {
	depth int
}

// NewASTPrinter creates a new instance of ASTPrinter and returns it.
func NewASTPrinter() *ASTPrinter {
	return &ASTPrinter{depth: 0}
}

// VisitProgram visits the Program node and prints it in a nice way
func (p *ASTPrinter) VisitProgram(program *ast.Program) {
	PrintIndentation(p.depth)
	println("Program:")
	p.depth++
	for _, stmt := range program.Stmts {
		if stmt != nil {
			stmt.Accept(p)
		} else {
			PrintNilIndentation(p.depth)
		}
	}
	p.depth--
}

// VisitStmtExit visits the StmtExit node and prints it in a nice way
func (p *ASTPrinter) VisitStmtExit(stmtExit *ast.StmtExit) {
	PrintIndentation(p.depth)
	println("StmtExit:")
	p.depth++
	if stmtExit.Expr != nil {
		stmtExit.Expr.Accept(p)
	} else {
		PrintNilIndentation(p.depth)
	}
	p.depth--
}

// VisitExprIntLit visits the ExprIntLit node and prints it in a nice way
func (p *ASTPrinter) VisitExprIntLit(exprIntLit *ast.ExprIntLit) {
	PrintIndentation(p.depth)
	println("ExprIntLit: ", exprIntLit.Val)
}
