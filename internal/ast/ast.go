// Package ast contains all the nodes and their definitions for tarka programming language.
package ast

// Node is the base interface for all tree elements
type Node interface {
	Print(depth int) // Print prints the current Node
}

// Expression is the base interface for all expression nodes
type Expression interface {
	Node
	exprNode()
}

// Statement is the base interface for all the statement nodes
type Statement interface {
	Node
	stmtNode()
}
