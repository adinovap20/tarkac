package ast

// ExprIntLit is the AST node structure for the integer literal expression
type ExprIntLit struct {
	Val  int // Val holds the value of the current integer literal
	Line int // Line number of the integer literal expression
	Col  int // Column number of the integer literal expression
}

func (e *ExprIntLit) exprNode() {}
func (e *ExprIntLit) Print(depth int) {
	PrintIndentation(depth)
	println("ExprIntLit: ", e.Val)
}
