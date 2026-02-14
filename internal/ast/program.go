package ast

// Program is the root of the Tarka AST
type Program struct {
	Stmts []Statement // Contains the list of statements of the Tarka program
}

func (p *Program) Print(depth int) {
	PrintIndentation(depth)
	println("Program:")
	for _, stmt := range p.Stmts {
		if stmt != nil {
			stmt.Print(depth + 1)
		} else {
			PrintNilIndentation(depth + 1)
		}
	}
}
