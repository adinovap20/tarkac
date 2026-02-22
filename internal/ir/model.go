// Package ir contains functions and structures for IR
package ir

import (
	"strconv"
	"strings"
)

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
	var str strings.Builder
	for _, inst := range p.Insts {
		str.WriteString(inst.String() + "\n")
	}
	return str.String()
}

// Accept visits the IRProgram
func (p *IRProgram) Accept(v Visitor) {
	v.VisitIRProgram(p)
}

// IRPushInt structure holds the IR for `LOAD INT <int>` instruction
type IRPushInt struct {
	Val int // Val holds the integer value to be loaded
}

// String returns the string representation of IRLoadInt
func (i *IRPushInt) String() string {
	return "PUSH INT " + strconv.Itoa(i.Val)
}

// Accept visits the IRLoadInt
func (i *IRPushInt) Accept(v Visitor) {
	v.VisitIRPushInt(i)
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

// IRStoreInt structure holds the IR for `STORE INT <identifier>`
type IRStoreInt struct {
	Name string // Name holds the identifier name
}

// String returns the string representation of IRStoreInt
func (i *IRStoreInt) String() string {
	return "STORE INT " + i.Name
}

// Accept visits the IRStoreInt
func (i *IRStoreInt) Accept(v Visitor) {
	v.VisitIRStoreInt(i)
}
