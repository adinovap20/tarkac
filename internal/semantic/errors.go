package semantic

import (
	"log"
	"strconv"
)

// newError adds the new error to the errors slice.
func (s *SemanticAnalyzer) newError(msg string, line, col int) {
	err := "Semantic Error: " + msg + ", at line " + strconv.Itoa(line) + ":" + strconv.Itoa(col)
	s.errors = append(s.errors, err)
}

// PrintErrors prints the erros on console
func (s *SemanticAnalyzer) PrintErrors() {
	for _, err := range s.errors {
		println(err)
	}
	if len(s.errors) > 0 {
		log.Fatalln("Semantic analysis failed! Errors occurred while parsing!")
	}
}
