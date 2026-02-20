package ir

// New creates and returns a new instance of IRProgram
func New() *IRProgram {
	return &IRProgram{}
}

// LoadInt adds a new IRLoadInt node for `LOAD INT <int>`
func (p *IRProgram) LoadInt(val int) {
	p.Insts = append(p.Insts, &IRLoadInt{Val: val})
}

// Exit adds a new IRExit node for `EXIT`
func (p *IRProgram) Exit() {
	p.Insts = append(p.Insts, &IRExit{})
}
