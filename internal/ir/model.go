// Package ir contains functions and structures for IR
package ir

import "strconv"

// IRInstruction interface for all the IR nodes
type IRInstruction interface {
	String() string // Returns the string representation
	Accept(v Visitor)
}

// IRProgram structure holds the IR as a program
type IRProgram struct {
	Insts []IRInstruction // List of IR Instructions
}

// String prints the IR program
func (p *IRProgram) String() string {
	str := ""
	for _, inst := range p.Insts {
		str += inst.String() + "\n"
	}
	return str
}

// Accept visits the IRProgram
func (p *IRProgram) Accept(v Visitor) {
	v.VisitIRProgram(p)
}

// IRLoadInt structure holds the IR for `LOAD INT <int>` instruction
type IRLoadInt struct {
	Val int // Val holds the integer value to be loaded
}

// String returns the string representation of IRLoadInt
func (i *IRLoadInt) String() string {
	return "LOAD INT " + strconv.Itoa(i.Val)
}

// Accept visits the IRLoadInt
func (i *IRLoadInt) Accept(v Visitor) {
	v.VisitIRLoadInt(i)
}

// IRExit structure holds the IR for `EXIT` instruction
type IRExit struct{}

// String returns the string representation of IRExit
func (i *IRExit) String() string {
	return "EXIT"
}

// Accept visits the IRExit
func (i *IRExit) Accept(v Visitor) {
	v.VisitIRExit(i)
}
