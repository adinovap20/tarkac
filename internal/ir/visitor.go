package ir

// Visitor interface for traversing the IR
type Visitor interface {
	VisitIRProgram(p *IRProgram)
	VisitIRPushInt(i *IRPushInt)
	VisitIRExit(i *IRExit)
	VisitIRStoreInt(i *IRStoreInt)
}
