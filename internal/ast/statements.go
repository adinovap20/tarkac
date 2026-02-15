package ast

// StmtExit is the AST node structure for the exit statement
type StmtExit struct {
	Expr Expression // Expr holds the expression that needs to be exited
	Line int        // Line number of the start of the exit statement
	Col  int        // Column number of the start of the exit statement
}

func (s *StmtExit) stmtNode() {}
func (s *StmtExit) Accept(v Visitor) {
	v.VisitStmtExit(s)
}
