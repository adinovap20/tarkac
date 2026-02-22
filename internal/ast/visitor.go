package ast

// Visitor interface for traversing the abstract syntax tree
type Visitor interface {
	VisitProgram(p *Program)
	VisitStmtExit(s *StmtExit)
	VisitStmtIntVarDecl(s *StmtIntVarDecl)
	VisitExprIntLit(e *ExprIntLit)
}
