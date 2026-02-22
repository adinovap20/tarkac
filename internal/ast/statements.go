package ast

// StmtIntVarDecl is the AST node structure for the int variable declaration
type StmtIntVarDecl struct {
	Name string     // Name of the variable being declared
	Expr Expression // Expr holds the expression that needs to be exited
	Line int        // Line number of the start of the variable declaration
	Col  int        // Column number of the start of the variable declaration
}

func (s *StmtIntVarDecl) stmtNode() {}
func (s *StmtIntVarDecl) Accept(v Visitor) {
	v.VisitStmtIntVarDecl(s)
}

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
