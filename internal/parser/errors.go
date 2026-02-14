package parser

import (
	"log"
	"strconv"
)

// newNoMatchingStmtError adds the no new matching stmt found error in the errors slice.
func (p *Parser) newNoMatchingStmtError() {
	err := "Parser Error: No matching statement found for token: '" + p.curToken.Lit + "', at line "
	err += strconv.Itoa(p.curToken.Line) + ":" + strconv.Itoa(p.curToken.Col)
	p.errors = append(p.errors, err)
}

// newNoMatchingExprError adds the no new matching expression found error in the errors slice.
func (p *Parser) newNoMatchingExprError() {
	err := "Parser Error: No matching expression found for token: '" + p.curToken.Lit + "', at line "
	err += strconv.Itoa(p.curToken.Line) + ":" + strconv.Itoa(p.curToken.Col)
	p.errors = append(p.errors, err)
}

// newError adds the new error to the errors slice.
func (p *Parser) newError(msg string, line, col int) {
	err := "Parser Error: " + msg + ", at line " + strconv.Itoa(line) + ":" + strconv.Itoa(col)
	p.errors = append(p.errors, err)
}

// newErrorWithoutLineStats adds the new error to the errors slice without line and column information
func (p *Parser) newErrorWithoutLineStats(msg string) {
	err := "Parser Error: " + msg
	p.errors = append(p.errors, err)
}

// PrintErrors prints the erros on console
func (p *Parser) PrintErrors() {
	for _, err := range p.errors {
		println(err)
	}
	if len(p.errors) > 0 {
		log.Fatalln("Syntax analysis failed! Errors occurred while parsing!")
	}
}
