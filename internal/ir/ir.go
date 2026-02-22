package ir

// New creates and returns a new instance of IRProgram
func New() *IRProgram {
	return &IRProgram{}
}

// PushInt adds a new IRLoadInt node for `LOAD INT <int>`
func (p *IRProgram) PushInt(val int) {
	p.Insts = append(p.Insts, &IRPushInt{Val: val})
}

// Exit adds a new IRExit node for `EXIT`
func (p *IRProgram) Exit() {
	p.Insts = append(p.Insts, &IRExit{})
}

// StoreInt adds a new IRStoreInt node for `STORE INT <identifier>`
func (p *IRProgram) StoreInt(name string) {
	p.Insts = append(p.Insts, &IRStoreInt{Name: name})
}
