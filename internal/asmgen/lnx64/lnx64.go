// Package lnx64 contains the code generation functions for IR to lnx64 FASM
package lnx64

import (
	"fmt"

	"github.com/adinovap20/tarkac/internal/ir"
)

// Generator structure for generating the FASM code
type Generator struct {
	Code string // Code holds the generated FASM code
}

// NewGenerator creates and returns a new instance of Generator
func NewGenerator() *Generator {
	return &Generator{Code: ""}
}

// addInst adds the instruction in the FASM code
func (g *Generator) addInst(inst string) {
	g.Code += inst + "\n"
}

// addIndentedInst adds the indented instruction in the FASM code.
func (g *Generator) addIndentedInst(inst string) {
	g.Code += "    " + inst + "\n"
}

// addBlankLine adds the blank line in the FASM code
func (g *Generator) addBlankLine() {
	g.Code += "\n"
}

// VisitIRProgram visits the IRProgram node and generates lnx64 FASM code.
func (g *Generator) VisitIRProgram(p *ir.IRProgram) {
	g.addInst("format ELF64 executable 3")
	g.addInst("entry start")
	g.addBlankLine()
	g.addInst("segment readable executable")
	g.addBlankLine()
	g.addInst("start:")
	for _, inst := range p.Insts {
		inst.Accept(g)
	}
}

// VisitIRLoadInt visits the IRLoadInt node and generates lnx64 FASM code.
func (g *Generator) VisitIRLoadInt(i *ir.IRLoadInt) {
	instComment := fmt.Sprintf("; --- LOAD INT %d ---", i.Val)
	inst := fmt.Sprintf("push %d", i.Val)
	g.addIndentedInst(instComment)
	g.addIndentedInst(inst)
	g.addBlankLine()
}

// VisitIRExit visits the IRExit node and generates lnx64 FASM code.
func (g *Generator) VisitIRExit(i *ir.IRExit) {
	instComment := "; --- EXIT ---"
	g.addIndentedInst(instComment)
	g.addIndentedInst("mov rax, 60")
	g.addIndentedInst("pop rdi")
	g.addIndentedInst("syscall")
	g.addBlankLine()
}
