// Package runner contains functions for running the compiler pipeline for compiling tarka source code
package runner

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/adinovap20/tarkac/internal/astprinter"
	"github.com/adinovap20/tarkac/internal/lexer"
	"github.com/adinovap20/tarkac/internal/parser"
	"github.com/adinovap20/tarkac/internal/token"
)

// Flags structure contains all the parsed data from the command line arguments or flags
type Flags struct {
	inputFile  *string // Pointer to the path of the string input file path
	outputFile *string // Pointer to the path of the string output file path
	debugFlag  *bool   // Pointer to the debug flag
}

// Run runs the compiler pipeline according to the command line options
func Run() {
	// Parse the command line arguments
	outputFile := flag.String("o", "out", "Output file path")
	debug := flag.Bool("d", false, "Enable debug mode prints")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Printf("Usage: tarkac <input_file> [options]\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	inputFile := args[0]
	cmdLineFlags := &Flags{inputFile: &inputFile, outputFile: outputFile, debugFlag: debug}
	if *cmdLineFlags.debugFlag {
		fmt.Printf("Compiling %s to %s with debug flag %v\n",
			*cmdLineFlags.inputFile, *cmdLineFlags.outputFile, *cmdLineFlags.debugFlag,
		)
	}

	// Run lexical analysis phas
	toks := runLexicalAnalysis(cmdLineFlags)
	runSyntaxAnalysis(cmdLineFlags, toks)
}

// runLexicalAnalysis runs the lexical analysis phase of the compiler pipeline
func runLexicalAnalysis(flags *Flags) []token.Token {
	if *flags.debugFlag {
		fmt.Println("=== Lexical Analysis ===")
	}
	content, err := os.ReadFile(*flags.inputFile)
	if err != nil {
		log.Fatal(err)
	}
	code := string(content)
	lexer := lexer.NewLexer(code)
	tokens := lexer.GetTokens()
	if *flags.debugFlag {
		token.PrintTokens(tokens)
		fmt.Println("Lexical analysis successful...")
	}
	return tokens
}

// runSyntaxAnalysis runs the syntax analysis phase of the compiler pipeline
func runSyntaxAnalysis(flags *Flags, tokens []token.Token) {
	if *flags.debugFlag {
		fmt.Println("=== Syntax Analysis ===")
	}
	parser := parser.NewParser(tokens)
	program := parser.Parse()
	if program == nil {
		log.Fatal("Syntax analysis failed... Parser returned <nil> abstract syntax tree!")
	}
	parser.PrintErrors()
	if *flags.debugFlag {
		printer := astprinter.NewASTPrinter()
		program.Accept(printer)
		fmt.Println("Syntax analysis successful...")
	}
}
