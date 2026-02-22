// Package lnx64 contains the code generation functions for IR to lnx64 FASM
package lnx64

import (
	"fmt"

	"github.com/adinovap20/tarkac/internal/ir"
)

// Generator structure for generating the FASM code
type Generator struct {
	Code      string // Code holds the generated FASM code
	offsets   map[string]int
	curOffset int
}

// NewGenerator creates and returns a new instance of Generator
func NewGenerator() *Generator {
	return &Generator{Code: "", offsets: map[string]int{}, curOffset: -8}
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
	g.addIndentedInst("; --- Prologue ---")
	g.addIndentedInst("push rbp")
	g.addIndentedInst("mov rbp, rsp")
	varSize := 0
	for _, inst := range p.Insts {
		if _, ok := inst.(*ir.IRStoreInt); ok {
			varSize += 8
		}
	}
	g.addIndentedInst(fmt.Sprintf("sub rsp, %d", varSize))
	g.addBlankLine()
	for _, inst := range p.Insts {
		inst.Accept(g)
	}
}

// VisitIRPushInt visits the IRPushInt node and generates lnx64 FASM code.
func (g *Generator) VisitIRPushInt(i *ir.IRPushInt) {
	instComment := fmt.Sprintf("; --- PUSH INT %d ---", i.Val)
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

// VisitIRStoreInt visits the IRStoreInt node and generates lnx64 FASM code.
func (g *Generator) VisitIRStoreInt(i *ir.IRStoreInt) {
	instComment := fmt.Sprintf("; --- STORE INT %s ---", i.Name)
	g.addIndentedInst(instComment)
	g.addIndentedInst("pop rax")
	var offset int
	if _, ok := g.offsets[i.Name]; !ok {
		g.offsets[i.Name] = g.curOffset
		offset = g.curOffset
		g.curOffset -= 8
	} else {
		offset = g.offsets[i.Name]
	}
	g.addIndentedInst(fmt.Sprintf("mov [rbp%d], rax", offset))
	g.addBlankLine()
}
