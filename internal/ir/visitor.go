package ir

// Visitor interface for traversing the IR
type Visitor interface {
	VisitIRProgram(p *IRProgram)
	VisitIRLoadInt(i *IRLoadInt)
	VisitIRExit(i *IRExit)
}
