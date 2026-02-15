package ast

// Program is the root of the Tarka AST
type Program struct {
	Stmts []Statement // Contains the list of statements of the Tarka program
}

func (p *Program) Accept(v Visitor) {
	v.VisitProgram(p)
}
