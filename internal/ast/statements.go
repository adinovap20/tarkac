package ast

// StmtExit is the AST node structure for the exit statement
type StmtExit struct {
	Expr Expression // Expr holds the expression that needs to be exited
	Line int        // Line number of the start of the exit statement
	Col  int        // Column number of the start of the exit statement
}

func (s *StmtExit) stmtNode() {}
func (s *StmtExit) Print(depth int) {
	PrintIndentation(depth)
	println("StmtExit:")
	if s.Expr != nil {
		s.Expr.Print(depth + 1)
	} else {
		PrintNilIndentation(depth + 1)
	}
}
